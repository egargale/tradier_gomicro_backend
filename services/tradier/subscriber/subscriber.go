package subscriber

import (
	"context"
	"log"

	"tradier_gomicro_backend/services/tradier/domain"
	proto  "tradier_gomicro_backend/services/tradier/proto"
)

var (
	Topic = "location"
)

type Location struct{}

func (g *Location) Handle(ctx context.Context, e *proto.StreamResponse) error {
	// log.Printf("Saving entity ID %s", e.Id)
	// domain.Save(ctx, domain.ProtoToEntity(e))
	return nil
}
