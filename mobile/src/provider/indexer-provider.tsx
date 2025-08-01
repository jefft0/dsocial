import { createContext, useContext, useEffect, useState } from "react";

import * as Grpc from "@gno/grpc/client";
import { Client } from "@connectrpc/connect";

import { HelloResponse, HelloStreamResponse, UserAndPostID } from "@buf/gnolang_dsocial-indexer.bufbuild_es/indexerservice_pb";
import { IndexerService } from "@buf/gnolang_dsocial-indexer.bufbuild_es/indexerservice_pb";

export interface IndexerContextProps {
  getHomePosts: (userPostAddr: string, startIndex: bigint, endIndex: bigint) => Promise<[number, string]>;
  hello: (name: string) => Promise<HelloResponse>;
  helloStream: (name: string) => Promise<AsyncIterable<HelloStreamResponse>>;
}

interface ConfigProps {
  remote: string;
}

interface IndexerProviderProps {
  config: ConfigProps;
  children: React.ReactNode;
}

const IndexerContext = createContext<IndexerContextProps | null>(null);

const IndexerProvider: React.FC<IndexerProviderProps> = ({ children, config }) => {
  const [clientInstance, setClientInstance] = useState<Client<typeof IndexerService> | undefined>(undefined);

  useEffect(() => {
    (async () => {
      if (clientInstance) {
        return; // Prevent re-initialization
      }
      setClientInstance(initClient(config));
    })();
  }, []);

  const initClient = (config: ConfigProps): Client<typeof IndexerService> => {
    if (clientInstance) {
      return clientInstance;
    }

    return Grpc.createIndexerClient(config.remote);
  };

  const getClient = () => {
    if (!clientInstance) {
      throw new Error("indexer client instance not initialized.");
    }

    return clientInstance;
  };

  const formatHomePost = (homePosts: UserAndPostID[]): string => {
    let result = "[]UserAndPostID{";
    for (const homePost of homePosts) {
      result += `{"${homePost.userPostAddr}", ${homePost.postID}},`;
    }
    result += "}";

    return result;
  };

  // Call getHomePosts and return [nHomePosts, addrAndIDs] where nHomePosts is the
  // total number of home posts and addrAndIDs is a Go string of the slice of
  // UserAndPostID which to use in qEval `GetJsonTopPostsByID(${addrAndIDs})`.
  const getHomePosts = async (userPostAddr: string, startIndex: bigint, endIndex: bigint): Promise<[number, string]> => {
    const client = getClient();

    const homePostsResult = await client.getHomePosts({
      userPostAddr,
      startIndex,
      endIndex,
    });
    const homePosts = formatHomePost(homePostsResult.homePosts);

    return [Number(homePostsResult.nPosts), homePosts];
  };

  const hello = async (name: string) => {
    const client = getClient();
    return client.hello({ name });
  };

  const helloStream = async (name: string) => {
    const client = getClient();
    return client.helloStream({ name });
  };

  if (!clientInstance) {
    return null;
  }

  const value = {
    getHomePosts,
    hello,
    helloStream,
  };

  return <IndexerContext.Provider value={value}>{children}</IndexerContext.Provider>;
};

function useIndexerContext() {
  const context = useContext(IndexerContext) as IndexerContextProps;

  if (context === undefined) {
    throw new Error("useIndexerContext must be used within a IndexerProvider");
  }
  return context;
}

export { IndexerProvider, useIndexerContext };
