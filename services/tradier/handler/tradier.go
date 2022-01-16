package handler

import (
	"context"
	"log"

	"tradier_gomicro_backend/services/tradier/domain"
	pb "tradier_gomicro_backend/services/tradier/proto"

	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
)

func (s *domain.Client) Stream(ctx context.Context, req *pb.StreamRequest, rsp *pb.StreamResponse) error {

	v1, err := config.Get("tradier.production.api")
	if err != nil {
		logger.Fatalf(" Tradier API config not found: %v", err)
	}

	v2, err := config.Get("tradier.production.account")
	if err != nil {
		logger.Fatalf("Tradier account not found: %v", err)
	}

	params := domain.DefaultParams(v1)
	client := domain.NewClient(params)
	client.SelectAccount(v2)

	_, session_err := client.GetSessionID()
	if session_err != nil {
		log.Printf("Opening Session: Error is %s", err)
	}

	return nil
}
