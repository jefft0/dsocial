// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: indexerservice.proto

package _goconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	_go "github.com/gnoverse/dsocial/tools/indexer-service/api/gen/go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// IndexerServiceName is the fully-qualified name of the IndexerService service.
	IndexerServiceName = "land.gno.dsocial.indexerservice.v1.IndexerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IndexerServiceGetHomePostsProcedure is the fully-qualified name of the IndexerService's
	// GetHomePosts RPC.
	IndexerServiceGetHomePostsProcedure = "/land.gno.dsocial.indexerservice.v1.IndexerService/GetHomePosts"
	// IndexerServiceStreamPostReplyProcedure is the fully-qualified name of the IndexerService's
	// StreamPostReply RPC.
	IndexerServiceStreamPostReplyProcedure = "/land.gno.dsocial.indexerservice.v1.IndexerService/StreamPostReply"
	// IndexerServiceHelloProcedure is the fully-qualified name of the IndexerService's Hello RPC.
	IndexerServiceHelloProcedure = "/land.gno.dsocial.indexerservice.v1.IndexerService/Hello"
	// IndexerServiceHelloStreamProcedure is the fully-qualified name of the IndexerService's
	// HelloStream RPC.
	IndexerServiceHelloStreamProcedure = "/land.gno.dsocial.indexerservice.v1.IndexerService/HelloStream"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	indexerServiceServiceDescriptor               = _go.File_indexerservice_proto.Services().ByName("IndexerService")
	indexerServiceGetHomePostsMethodDescriptor    = indexerServiceServiceDescriptor.Methods().ByName("GetHomePosts")
	indexerServiceStreamPostReplyMethodDescriptor = indexerServiceServiceDescriptor.Methods().ByName("StreamPostReply")
	indexerServiceHelloMethodDescriptor           = indexerServiceServiceDescriptor.Methods().ByName("Hello")
	indexerServiceHelloStreamMethodDescriptor     = indexerServiceServiceDescriptor.Methods().ByName("HelloStream")
)

// IndexerServiceClient is a client for the land.gno.dsocial.indexerservice.v1.IndexerService
// service.
type IndexerServiceClient interface {
	GetHomePosts(context.Context, *connect.Request[_go.GetHomePostsRequest]) (*connect.Response[_go.GetHomePostsResponse], error)
	// StreamPostReply returns a stream of PostReply functions that are called in the blockchain
	StreamPostReply(context.Context, *connect.Request[_go.StreamPostReplyRequest]) (*connect.ServerStreamForClient[_go.StreamPostReplyResponse], error)
	// Hello is for debug purposes
	Hello(context.Context, *connect.Request[_go.HelloRequest]) (*connect.Response[_go.HelloResponse], error)
	// HelloStream is for debug purposes
	HelloStream(context.Context, *connect.Request[_go.HelloStreamRequest]) (*connect.ServerStreamForClient[_go.HelloStreamResponse], error)
}

// NewIndexerServiceClient constructs a client for the
// land.gno.dsocial.indexerservice.v1.IndexerService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIndexerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IndexerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &indexerServiceClient{
		getHomePosts: connect.NewClient[_go.GetHomePostsRequest, _go.GetHomePostsResponse](
			httpClient,
			baseURL+IndexerServiceGetHomePostsProcedure,
			connect.WithSchema(indexerServiceGetHomePostsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		streamPostReply: connect.NewClient[_go.StreamPostReplyRequest, _go.StreamPostReplyResponse](
			httpClient,
			baseURL+IndexerServiceStreamPostReplyProcedure,
			connect.WithSchema(indexerServiceStreamPostReplyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		hello: connect.NewClient[_go.HelloRequest, _go.HelloResponse](
			httpClient,
			baseURL+IndexerServiceHelloProcedure,
			connect.WithSchema(indexerServiceHelloMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		helloStream: connect.NewClient[_go.HelloStreamRequest, _go.HelloStreamResponse](
			httpClient,
			baseURL+IndexerServiceHelloStreamProcedure,
			connect.WithSchema(indexerServiceHelloStreamMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// indexerServiceClient implements IndexerServiceClient.
type indexerServiceClient struct {
	getHomePosts    *connect.Client[_go.GetHomePostsRequest, _go.GetHomePostsResponse]
	streamPostReply *connect.Client[_go.StreamPostReplyRequest, _go.StreamPostReplyResponse]
	hello           *connect.Client[_go.HelloRequest, _go.HelloResponse]
	helloStream     *connect.Client[_go.HelloStreamRequest, _go.HelloStreamResponse]
}

// GetHomePosts calls land.gno.dsocial.indexerservice.v1.IndexerService.GetHomePosts.
func (c *indexerServiceClient) GetHomePosts(ctx context.Context, req *connect.Request[_go.GetHomePostsRequest]) (*connect.Response[_go.GetHomePostsResponse], error) {
	return c.getHomePosts.CallUnary(ctx, req)
}

// StreamPostReply calls land.gno.dsocial.indexerservice.v1.IndexerService.StreamPostReply.
func (c *indexerServiceClient) StreamPostReply(ctx context.Context, req *connect.Request[_go.StreamPostReplyRequest]) (*connect.ServerStreamForClient[_go.StreamPostReplyResponse], error) {
	return c.streamPostReply.CallServerStream(ctx, req)
}

// Hello calls land.gno.dsocial.indexerservice.v1.IndexerService.Hello.
func (c *indexerServiceClient) Hello(ctx context.Context, req *connect.Request[_go.HelloRequest]) (*connect.Response[_go.HelloResponse], error) {
	return c.hello.CallUnary(ctx, req)
}

// HelloStream calls land.gno.dsocial.indexerservice.v1.IndexerService.HelloStream.
func (c *indexerServiceClient) HelloStream(ctx context.Context, req *connect.Request[_go.HelloStreamRequest]) (*connect.ServerStreamForClient[_go.HelloStreamResponse], error) {
	return c.helloStream.CallServerStream(ctx, req)
}

// IndexerServiceHandler is an implementation of the
// land.gno.dsocial.indexerservice.v1.IndexerService service.
type IndexerServiceHandler interface {
	GetHomePosts(context.Context, *connect.Request[_go.GetHomePostsRequest]) (*connect.Response[_go.GetHomePostsResponse], error)
	// StreamPostReply returns a stream of PostReply functions that are called in the blockchain
	StreamPostReply(context.Context, *connect.Request[_go.StreamPostReplyRequest], *connect.ServerStream[_go.StreamPostReplyResponse]) error
	// Hello is for debug purposes
	Hello(context.Context, *connect.Request[_go.HelloRequest]) (*connect.Response[_go.HelloResponse], error)
	// HelloStream is for debug purposes
	HelloStream(context.Context, *connect.Request[_go.HelloStreamRequest], *connect.ServerStream[_go.HelloStreamResponse]) error
}

// NewIndexerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIndexerServiceHandler(svc IndexerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	indexerServiceGetHomePostsHandler := connect.NewUnaryHandler(
		IndexerServiceGetHomePostsProcedure,
		svc.GetHomePosts,
		connect.WithSchema(indexerServiceGetHomePostsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	indexerServiceStreamPostReplyHandler := connect.NewServerStreamHandler(
		IndexerServiceStreamPostReplyProcedure,
		svc.StreamPostReply,
		connect.WithSchema(indexerServiceStreamPostReplyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	indexerServiceHelloHandler := connect.NewUnaryHandler(
		IndexerServiceHelloProcedure,
		svc.Hello,
		connect.WithSchema(indexerServiceHelloMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	indexerServiceHelloStreamHandler := connect.NewServerStreamHandler(
		IndexerServiceHelloStreamProcedure,
		svc.HelloStream,
		connect.WithSchema(indexerServiceHelloStreamMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/land.gno.dsocial.indexerservice.v1.IndexerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IndexerServiceGetHomePostsProcedure:
			indexerServiceGetHomePostsHandler.ServeHTTP(w, r)
		case IndexerServiceStreamPostReplyProcedure:
			indexerServiceStreamPostReplyHandler.ServeHTTP(w, r)
		case IndexerServiceHelloProcedure:
			indexerServiceHelloHandler.ServeHTTP(w, r)
		case IndexerServiceHelloStreamProcedure:
			indexerServiceHelloStreamHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIndexerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIndexerServiceHandler struct{}

func (UnimplementedIndexerServiceHandler) GetHomePosts(context.Context, *connect.Request[_go.GetHomePostsRequest]) (*connect.Response[_go.GetHomePostsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.dsocial.indexerservice.v1.IndexerService.GetHomePosts is not implemented"))
}

func (UnimplementedIndexerServiceHandler) StreamPostReply(context.Context, *connect.Request[_go.StreamPostReplyRequest], *connect.ServerStream[_go.StreamPostReplyResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.dsocial.indexerservice.v1.IndexerService.StreamPostReply is not implemented"))
}

func (UnimplementedIndexerServiceHandler) Hello(context.Context, *connect.Request[_go.HelloRequest]) (*connect.Response[_go.HelloResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.dsocial.indexerservice.v1.IndexerService.Hello is not implemented"))
}

func (UnimplementedIndexerServiceHandler) HelloStream(context.Context, *connect.Request[_go.HelloStreamRequest], *connect.ServerStream[_go.HelloStreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.dsocial.indexerservice.v1.IndexerService.HelloStream is not implemented"))
}
