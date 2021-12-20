package main

import (
	"tests/handler"
	"tests/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tradier"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTradierHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
