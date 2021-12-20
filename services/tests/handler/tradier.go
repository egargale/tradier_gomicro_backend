package handler

import (
	"context"
	"time"

	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"
	tradierlib "github.com/timpalpant/go-tradier"

	pb "tradier/proto"
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

// New
func New() *Tradier {
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

func (s *Tradier) Quote(ctx context.Context, req *pb.QuoteRequest, rsp *pb.QuoteResponse) error {
	if len(req.Symbol) <= 0 {
		return errors.BadRequest("Tradier.quote", "invalid symbol")
	}
	params := tradierlib.DefaultParams(s.Api)
	client := tradierlib.NewClient(params)

	// uri := s.Endpoint + "markets/quotes"
	// uri := fmt.Sprintf("%slast/forex/%s?apikey=%s", s.Api, req.Symbol, s.Key)

	//quotes, err := client.GetQuotes([]string{"AAPL", "SPY"})
	quotes_res, err := client.GetQuotes([]string{req.Symbol})
	if err != nil {
		logger.Errorf("Failed to get quote: %v\n", err)
		return errors.InternalServerError("tradier.quote", "failed to get quote")
	}
	for _, quote := range quotes_res {
		info := &pb.QuoteInfo{
			Symbol:   quote.Symbol,
			Last:     quote.Last,
			BidPrice: quote.Bid,
			BidSize:  int32(quote.BidSize),
			AskPrice: quote.Ask,
			AskSize:  int32(quote.AskSize),
		}
		rsp.Quote = append(rsp.Quote, info)
	}

	return nil
}
func (s *Tradier) GetMarketState(ctx context.Context, req *pb.MarketeStateRequest, rsp *pb.MarketeStateResponse) error {

	params := tradierlib.DefaultParams(s.Api)
	// params := tradierlib.DefaultParams("hmEh5bdeZgl9tq6QpuEIVfRNQPlb")
	client := tradierlib.NewClient(params)

	clock, err := client.GetMarketState()
	if err != nil {
		logger.Errorf("Failed to get market state: %v\n", err)
		return errors.InternalServerError("tradier.getmarketstate", "failed to get market state")
	}
	//rsp.Clock = fmt.Sprintf("%#v", clock)
	rsp.Time = clock.Time.String()
	rsp.State = clock.State
	rsp.Description = clock.Description
	rsp.Nextstate = clock.NextState

	return nil
}
