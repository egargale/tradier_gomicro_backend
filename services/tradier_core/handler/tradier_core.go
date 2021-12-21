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
