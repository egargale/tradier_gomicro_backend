package main

import (
	"mytest/handler"
	pb "github.com/egargale/tradier_gomicro_backend/services/mytest"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("mytest"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMytestHandler(srv.Server(), new(handler.Mytest))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
