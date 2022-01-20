package main

import (
	"context"
	"encoding/json"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/broker"
	"github.com/micro/micro/v3/service/logger"
	"github.com/timpalpant/go-tradier"

	pb "github.com/egargale/tradier_gomicro_backend/services/test2/proto"
)

type Healthcheck struct {
	Healthy      bool
	Service      string
	Notification string
}

type Market struct {
	Time        tradier.DateTime
	State       string
	Description string
	NextChange  tradier.DateTime
	NextState   string
}

func main() {

	srv := service.New()

	handler := func(msg *broker.Message) error {
		var hc Healthcheck
		if err := json.Unmarshal(msg.Body, &hc); err != nil {
			return err
		}

		if hc.Healthy {
			logger.Infof("Service %v healthy and notification is %s", hc.Service, hc.Notification)
		} else {
			logger.Infof("Service %v is not healthy", hc.Service)
		}

		return nil
	}

	handler_market := func(msg *broker.Message) error {
		var mc Market
		if err := json.Unmarshal(msg.Body, &mc); err != nil {
			return err
		}
		// logger.Infof("Time: %s State: %s %s %s", mc.Time.String(), mc.State, mc.Description, mc.NextChange.String())
		logger.Infof(string(msg.Body))
		if mc.State == "open" {
			srv := service.New()

			// create the proto client for helloworld
			client := pb.NewTest2Service("test2", srv.Client())

			// call an endpoint on the service
			rsp, err := client.Call(context.Background(), &pb.Request{
				Name: "Market",
			})
			if err != nil {
				logger.Errorf("Error calling test2: ", err)
				return err
			}
			logger.Infof("From test2 service response is %s", rsp.Msg)
		}
		return nil
	}

	_, err := broker.Subscribe("health", handler)
	if err != nil {
		logger.Fatal(err)
	}

	_, err1 := broker.Subscribe("market", handler_market)
	if err1 != nil {
		logger.Fatal(err1)
	}

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
