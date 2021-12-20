package main

import (
	"tradier_core/handler"
	pb "tradier_core/proto"

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
	pb.RegisterTradier_coreHandler(srv.Server(), new(handler.Tradier_core))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
