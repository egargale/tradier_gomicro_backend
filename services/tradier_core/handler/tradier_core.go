package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	tradier_core "github.com/egargale/tradier_gomicro_backend/services/tradier_core/proto"
)

type Tradier_core struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Tradier_core) Call(ctx context.Context, req *tradier_core.Request, rsp *tradier_core.Response) error {
	log.Info("Received Tradier_core.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Tradier_core) Stream(ctx context.Context, req *tradier_core.StreamingRequest, stream tradier_core.Tradier_core_StreamStream) error {
	log.Infof("Received Tradier_core.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&tradier_core.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Tradier_core) PingPong(ctx context.Context, stream tradier_core.Tradier_core_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&tradier_core.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
