package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	test2 "test2/proto"
)

type Test2 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test2) Call(ctx context.Context, req *test2.Request, rsp *test2.Response) error {
	log.Info("Received Test2.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test2) Stream(ctx context.Context, req *test2.StreamingRequest, stream test2.Test2_StreamStream) error {
	log.Infof("Received Test2.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&test2.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test2) PingPong(ctx context.Context, stream test2.Test2_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&test2.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
