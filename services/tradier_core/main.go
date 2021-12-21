package main

import (
	"github.com/egargale/tradier_gomicro_backend/services/tradier_core/handler"
	pb "github.com/egargale/tradier_gomicro_backend/services/tradier_core/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tradier_core"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTradierCoreHandler(srv.Server(), new(handler.Tradier_core))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
