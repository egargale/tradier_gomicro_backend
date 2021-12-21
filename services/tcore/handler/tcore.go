package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	tcore "tcore/proto"
)

type Tcore struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Tcore) Call(ctx context.Context, req *tcore.Request, rsp *tcore.Response) error {
	log.Info("Received Tcore.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Tcore) Stream(ctx context.Context, req *tcore.StreamingRequest, stream tcore.Tcore_StreamStream) error {
	log.Infof("Received Tcore.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&tcore.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Tcore) PingPong(ctx context.Context, stream tcore.Tcore_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&tcore.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
