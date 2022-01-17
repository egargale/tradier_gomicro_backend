package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	mytest "github.com/egargale/tradier_gomicro_backend/services/mytest"
)

type Mytest struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Mytest) Call(ctx context.Context, req *mytest.Request, rsp *mytest.Response) error {
	log.Info("Received Mytest.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Mytest) Stream(ctx context.Context, req *mytest.StreamingRequest, stream mytest.Mytest_StreamStream) error {
	log.Infof("Received Mytest.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&mytest.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Mytest) PingPong(ctx context.Context, stream mytest.Mytest_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&mytest.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
