package main

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"go.uber.org/zap"

	api_gen "github.com/gnoverse/dsocial/tools/indexer-service/api/gen/go"
)

func (s *indexerService) GetHomePosts(ctx context.Context, req *connect.Request[api_gen.GetHomePostsRequest]) (*connect.Response[api_gen.GetHomePostsResponse], error) {
	nPosts := 0
	data := []*api_gen.UserAndPostID{}
	userPosts, ok := gUserPostsByAddress[req.Msg.UserPostAddr]
	if ok {
		nPosts = len(userPosts.homePosts)
		for i := int(req.Msg.StartIndex); i < int(req.Msg.EndIndex) && i < len(userPosts.homePosts); i++ {
			data = append(data, &api_gen.UserAndPostID{
				UserPostAddr: userPosts.homePosts[i].UserPostAddr,
				PostID:       uint64(userPosts.homePosts[i].PostID),
			})
		}
	}

	return connect.NewResponse(&api_gen.GetHomePostsResponse{
		NPosts:    uint64(nPosts),
		HomePosts: data,
	}), nil
}

// StreamPostReply return a stream of post replies
func (s *indexerService) StreamPostReply(ctx context.Context, req *connect.Request[api_gen.StreamPostReplyRequest], stream *connect.ServerStream[api_gen.StreamPostReplyResponse]) error {
	s.logger.Debug("StreamPostReply called")

	for {
		select {
		case <-ctx.Done():
			s.logger.Debug("StreamPostReply context done", zap.String("ctx.Err()", ctx.Err().Error()))
			return nil
		case msg := <-s.cPostReply:
			if err := stream.Send(msg); err != nil {
				s.logger.Error("StreamPostReply returned error", zap.Error(err))
				return err
			}
		}
	}
}

// Hello is for debug purposes
func (s *indexerService) Hello(ctx context.Context, req *connect.Request[api_gen.HelloRequest]) (*connect.Response[api_gen.HelloResponse], error) {
	s.logger.Debug("Hello called")
	defer s.logger.Debug("Hello returned ok")
	return connect.NewResponse(&api_gen.HelloResponse{
		Greeting: "Hello " + req.Msg.Name,
	}), nil
}

// HelloStream is for debug purposes
func (s *indexerService) HelloStream(ctx context.Context, req *connect.Request[api_gen.HelloStreamRequest], stream *connect.ServerStream[api_gen.HelloStreamResponse]) error {
	s.logger.Debug("HelloStream called")
	for i := 0; i < 4; i++ {
		if err := stream.Send(&api_gen.HelloStreamResponse{
			Greeting: "Hello " + req.Msg.Name,
		}); err != nil {
			s.logger.Error("HelloStream returned error", zap.Error(err))
			return err
		}
		time.Sleep(2 * time.Second)
	}

	s.logger.Debug("HelloStream returned ok")
	return nil
}
