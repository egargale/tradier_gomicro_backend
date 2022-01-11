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
func Setup() *Tradier {
	//
	v, err := config.Get("tradier.sandbox.api")
	if err != nil {
		logger.Fatalf(" Tradier API config not found: %v", err)
	}
	api := v.String("")
	if len(api) == 0 {
		logger.Fatal("Tardier API not found")
	}
	v, err = config.Get("tradier.sandbox.account")
	if err != nil {
		logger.Fatalf("Tradier account not found: %v", err)
	}
	account := v.String("")
	if len(account) == 0 {
		logger.Fatal("Tradier account not found")
	}
	v, err = config.Get("tradier.sandbox.endpoint")
	if err != nil {
		logger.Fatalf("Tradier endpoint not found: %v", err)
	}
	endpoint := v.String("")
	if len(endpoint) == 0 {
		logger.Fatal("Tradier endpoint not found")
	}
	return &Tradier{
		Api:      api,
		Account:  account,
		Endpoint: endpoint,
	}
}
