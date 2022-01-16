package domain

import (
	"time"

	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
)

type DateTime struct {
	time.Time
}

type Tradier struct {
	Account  string
	Api      string
	Endpoint string
}

type MarketStatus struct {
	Time        DateTime `json:"date"`
	State       string
	Description string
	//NextChange  DateTime `json:"next_change"`
	NextState string `json:"next_state"`
}

// Setup
func Setup() *ClientParams {
	//
	v, err := config.Get("tradier.production.api")
	if err != nil {
		logger.Fatalf(" Tradier API config not found: %v", err)
	}
	api := v.String("")
	if len(api) == 0 {
		logger.Fatal("Tardier API not found")
	}
	v, err = config.Get("tradier.production.account")
	if err != nil {
		logger.Fatalf("Tradier account not found: %v", err)
	}
	account := v.String("")
	if len(account) == 0 {
		logger.Fatal("Tradier account not found")
	}
	return &ClientParams{
		AuthToken: api,
		Account:   account,
	}
}
