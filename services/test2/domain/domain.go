package domain

import (
	// "time"
	"encoding/json"

	"github.com/micro/micro/v3/service/broker"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	tradier "github.com/timpalpant/go-tradier"
)

// Setup
func Setup() *tradier.ClientParams {
	//
	// v, err := config.Get("tradier.production.api")
	// if err != nil {
	// 	logger.Fatalf(" Tradier API config not found: %v", err)
	// }
	// api := v.String("")
	// if len(api) == 0 {
	// 	logger.Fatal("Tardier API not found")
	// }
	// v, err = config.Get("tradier.production.account")
	// if err != nil {
	// 	logger.Fatalf("Tradier account not found: %v", err)
	// }
	// account := v.String("")
	// if len(account) == 0 {
	// 	logger.Fatal("Tradier account not found")
	// }
	// return &ClientParams{
	// 	AuthToken: api,
	// 	Account:   account,
	// }
	return nil
}

func Status() {
	v, err := config.Get("tradier.production.api")
	if err != nil {
		logger.Fatalf(" Tradier API config not found: %v", err)
	}
	api := v.String("")
	if len(api) == 0 {
		logger.Fatal("Tardier API not found")
	}
	params := tradier.DefaultParams(api)
	client := tradier.NewClient(params)

	market, err := client.GetMarketState()
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(&tradier.MarketStatus{
		Time:      	  market.Time,
		State:     	  market.State,
		Description:  market.Description,
		NextChange:   market.NextChange,
		NextState:    market.NextState,
	})

	if err != nil {
		logger.Errorf("Error publishing ", err)
		return
	}

	broker.Publish("health", &broker.Message{Body: bytes})
}
