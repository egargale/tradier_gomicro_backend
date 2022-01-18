package main

import (
	"encoding/json"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/broker"
	"github.com/micro/micro/v3/service/logger"
)

type Healthcheck struct {
	Healthy      bool
	Service      string
	Notification string
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

	_, err := broker.Subscribe("health", handler)
	if err != nil {
		logger.Fatal(err)
	}

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
