package main

import (
	"tradier_gomicro_backend/services/tests/handler"
	pb "tradier_gomicro_backend/services/tests/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tests"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTestsHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
