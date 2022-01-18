package main

import (
	"tstream/handler"
	pb "tstream/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tstream"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTstreamHandler(srv.Server(), new(handler.Tstream))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
