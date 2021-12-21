package main

import (
	"tcore/handler"
	pb "github.com/egargale/tradier_gomicro_backend/services/tcore/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tcore"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTcoreHandler(srv.Server(), new(handler.Tcore))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
