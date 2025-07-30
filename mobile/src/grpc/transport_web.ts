// @ts-ignore
import { polyfill as polyfillReadableStream } from "react-native-polyfill-globals/src/readable-stream";
polyfillReadableStream();
// @ts-ignore
import { fetch as fetchPolyfill, Headers as HeadersPolyfill } from "react-native-fetch-api";

import type {
  DescMessage,
  DescMethodUnary,
  DescMethodStreaming,
  MessageShape,
  MessageInitShape,
} from '@bufbuild/protobuf';
import { toJsonString } from '@bufbuild/protobuf';
import type { ContextValues, StreamResponse, Transport, UnaryRequest, UnaryResponse } from "@connectrpc/connect";
import { createContextValues } from "@connectrpc/connect";
import {
  createClientMethodSerializers,
  createEnvelopeReadableStream,
  createMethodUrl,
  encodeEnvelope,
  runStreamingCall,
  runUnaryCall,
} from "@connectrpc/connect/protocol";
import {
  endStreamFlag,
  endStreamFromJson,
  validateResponse,
  requestHeader,
} from "@connectrpc/connect/protocol-connect";
import {
  trailerFlag,
  trailerParse,
  validateResponse as webValidateResponse,
  validateTrailer,
} from "@connectrpc/connect/protocol-grpc-web";
import { GrpcWebTransportOptions } from "@connectrpc/connect-web";

// Polyfill async.Iterator. For some reason, the Babel presets and plugins are not doing the trick.
// Code from here: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-3.html#caveats
(Symbol as any).asyncIterator = Symbol.asyncIterator || Symbol.for("Symbol.asyncIterator");

class AbortError extends Error {
  name = "AbortError";
}

interface FetchXHRResponse {
  status: number;
  headers: Headers;
  body: Uint8Array;
}

const fetchOptions: RequestInit = {
  redirect: "error",
};

function parseHeaders(allHeaders: string): Headers {
  return allHeaders
    .trim()
    .split(/[\r\n]+/)
    .reduce((memo, header) => {
      const [key, value] = header.split(": ");
      memo.append(key, value);
      return memo;
    }, new Headers());
}

function extractDataChunks(initialData: Uint8Array) {
  let buffer = initialData;
  let dataChunks: { flags: number; data: Uint8Array }[] = [];

  while (buffer.byteLength >= 5) {
    let length = 0;
    let flags = buffer[0];

    for (let i = 1; i < 5; i++) {
      length = (length << 8) + buffer[i]; // eslint-disable-line no-bitwise
    }

    const data = buffer.subarray(5, 5 + length);
    buffer = buffer.subarray(5 + length);
    dataChunks.push({ flags, data });
  }

  return dataChunks;
}

export function createXHRGrpcWebTransport(options: GrpcWebTransportOptions): Transport {
  const useBinaryFormat = options.useBinaryFormat ?? true;
  return {
    async unary<I extends DescMessage, O extends DescMessage>(
      method: DescMethodUnary<I, O>,
      signal: AbortSignal | undefined,
      timeoutMs: number | undefined,
      header: Headers,
      message: MessageInitShape<I>,
      contextValues?: ContextValues
    ): Promise<UnaryResponse<I, O>> {
      const { serialize, parse } = createClientMethodSerializers(
        method,
        useBinaryFormat,
        options.jsonOptions,
        options.binaryOptions
      );

      return await runUnaryCall<I, O>({
        signal,
        interceptors: options.interceptors,
        req: {
          stream: false,
          service: method.parent,
          method,
          requestMethod: 'POST',
          url: createMethodUrl(options.baseUrl, method),
          header: requestHeader(method.methodKind, useBinaryFormat, timeoutMs, header, false),
          contextValues: contextValues ?? createContextValues(),
          message,
        },
        next: async (req: UnaryRequest<I, O>): Promise<UnaryResponse<I, O>> => {
          function fetchXHR(): Promise<FetchXHRResponse> {
            return new Promise((resolve, reject) => {
              const xhr = new XMLHttpRequest();

              xhr.open(req.requestMethod ?? "POST", req.url);

              function onAbort() {
                xhr.abort();
              }

              req.signal.addEventListener("abort", onAbort);

              xhr.addEventListener("abort", () => {
                reject(new AbortError("Request aborted"));
              });

              xhr.addEventListener("load", () => {
                resolve({
                  status: xhr.status,
                  headers: parseHeaders(xhr.getAllResponseHeaders()),
                  body: new Uint8Array(xhr.response),
                });
              });

              xhr.addEventListener("error", () => {
                reject(new Error("Network Error"));
              });

              xhr.addEventListener("loadend", () => {
                req.signal.removeEventListener("abort", onAbort);
              });

              xhr.responseType = "arraybuffer";

              req.header.forEach((value: string, key: string) => xhr.setRequestHeader(key, value));

              xhr.send(encodeEnvelope(0, serialize(req.message)));
            });
          }
          const response = await fetchXHR();

          webValidateResponse(response.status, response.headers);

          const chunks = extractDataChunks(response.body);

          let trailer: Headers | undefined;
          let message: MessageShape<O> | undefined;

          chunks.forEach(({ flags, data }) => {
            if (flags === trailerFlag) {
              if (trailer !== undefined) {
                throw "extra trailer";
              }

              // Unary responses require exactly one response message, but in
              // case of an error, it is perfectly valid to have a response body
              // that only contains error trailers.
              trailer = trailerParse(data);
              return;
            }

            if (message !== undefined) {
              throw "extra message";
            }

            message = parse(data);
          });

          if (trailer === undefined) {
            throw "missing trailer";
          }

          validateTrailer(trailer, response.headers);

          if (message === undefined) {
            throw "missing message";
          }

          return <UnaryResponse<I, O>>{
            stream: false,
            header: response.headers,
            message,
            trailer,
          };
        },
      });
    },

    async stream<I extends DescMessage, O extends DescMessage>(
      method: DescMethodStreaming<I, O>,
      signal: AbortSignal | undefined,
      timeoutMs: number | undefined,
      header: HeadersInit | undefined,
      input: AsyncIterable<MessageInitShape<I>>,
      contextValues?: ContextValues
    ): Promise<StreamResponse<I, O>> {
      const { serialize, parse } = createClientMethodSerializers(
        method,
        useBinaryFormat,
        options.jsonOptions,
        options.binaryOptions
      );

      async function* parseResponseBody(body: ReadableStream<Uint8Array>, trailerTarget: HeadersPolyfill) {
        const reader = createEnvelopeReadableStream(body).getReader();
        let endStreamReceived = false;

        for (;;) {
          const result = await reader.read();
          if (result.done) {
            break;
          }

          const { flags, data } = result.value;
          if ((flags & endStreamFlag) === endStreamFlag) {
            endStreamReceived = true;

            const endStream = endStreamFromJson(data);
            if (endStream.error) {
              throw endStream.error;
            }

            endStream.metadata.forEach((value:any, key:any) => trailerTarget.set(key, value));
            continue;
          }

          yield parse(data);
        }

        if (!endStreamReceived) {
          throw "missing EndStreamResponse";
        }
      }

      async function createRequestBody(
        input: AsyncIterable<MessageShape<I>>,
      ): Promise<Uint8Array> {
        if (method.methodKind !== 'server_streaming') {
          throw 'The fetch API does not support streaming request bodies';
        }

        const r = await input[Symbol.asyncIterator]().next();
        if (r.done == true) {
          throw "missing request message";
        }

        return encodeEnvelope(0, serialize(r.value));
      }

      return await runStreamingCall<I, O>({
        interceptors: options.interceptors,
        timeoutMs,
        signal,
        req: {
          stream: true,
          service: method.parent,
          method,
          requestMethod: 'POST',
          url: createMethodUrl(options.baseUrl, method),
          header: requestHeader(method.methodKind, useBinaryFormat, timeoutMs, header, false),
          contextValues: contextValues ?? createContextValues(),
          message: input,
        },
        next: async (req) => {
          const fRes = await fetchPolyfill(req.url, {
            ...fetchOptions,
            headers: req.header,
            signal: req.signal,
            body: await createRequestBody(req.message),
            reactNative: { textStreaming: true }, // allows streaming in the polyfill fetch function
          });

          validateResponse(method.methodKind, useBinaryFormat, fRes.status, fRes.headers);
          if (fRes.body === null) {
            throw "missing response body";
          }

          const trailer = new HeadersPolyfill();

          // We have to implement the `*[Symbol.asyncIterator]()` of the object we give to the StreamResponse.message field.
          // It seems that react-native lacks the feature.
          const generator = {
            async *[Symbol.asyncIterator]() {
              yield* parseResponseBody(fRes.body, trailer);
            },
          };

          const res: StreamResponse<I, O> = {
            ...req,
            header: fRes.headers,
            trailer,
            message: generator,
          };
          return res;
        },
      });
    },
  };
}
