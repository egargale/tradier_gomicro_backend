package handler

import (
	"context"

	"tstream/domain"
	tstream "tstream/proto"

	log "github.com/micro/micro/v3/service/logger"
)

type Tstream struct{}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Tstream) Stream(ctx context.Context, req *tstream.StreamingRequest, stream tstream.Tstream_StreamStream) error {
	domain.Setup()
	domain.Start()

	log.Infof("Received Tstream.Stream request with count: %s", req.Count)

	for i := 0; i < 100; i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&tstream.StreamingResponse{
			Reply: "Hello World",
		}); err != nil {
			return err
		}
	}

	return nil
}
