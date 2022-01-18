package main

import (
	"github.com/egargale/tradier_gomicro_backend.git/services/test2/handler"
	pb "github.com/egargale/tradier_gomicro_backend.git/services/test2/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("test2"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTest2Handler(srv.Server(), new(handler.Test2))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
