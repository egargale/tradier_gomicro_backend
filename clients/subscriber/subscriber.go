package main

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"strconv"

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
func publishStream(service *service.Service, resp pb.StreamingResponse) error {

	bytes, err := json.Marshal(&Healthcheck{
		Healthy:      true,
		Service:      service.Options().Name,
		Notification: strconv.Itoa(int(resp.Count)),
	})

	if err != nil {
		logger.Errorf("Error publishing ", err)
		return err
	}

	stream_err := broker.Publish("health", &broker.Message{Body: bytes})
	if err != nil {
		logger.Errorf("Message not brocasted: Error", stream_err)
		return stream_err
	}
	return nil
} 

func getHelloMarket(service *service.Service, proto pb.Test2Service) {

	// call an endpoint on the service
	rsp, err := proto.Call(context.Background(), &pb.Request{
		Name: "Market",
	})
	if err != nil {
		logger.Errorf("Error calling test2: ", err)
		return
	}
	logger.Infof("From test2 service response is %s", rsp.Msg)
}

func serverStream(service *service.Service, proto pb.Test2Service) {
	// send request to stream count of 10
	stream, err := proto.Stream(context.Background(), &pb.StreamingRequest{Count: int64((rand.Intn(100)))})
	if err != nil {
		logger.Errorf("err:", err)
		return
	}
	defer stream.Close()

	// server side stream
	// receive messages for a random countdown
	for {
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Fatal(err)
		}
		// Publish to broker
		pub_err := publishStream(service , *rsp)
		if pub_err != nil {
			logger.Errorf("Error in broadcasting: Error", pub_err)
		}
		logger.Infof("got msg %v\n", rsp.Count)
	}
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
		switch mc.State {
		case "open":
			// create and initialise a new service
			srv := service.New()
			// create the proto client for
			client := pb.NewTest2Service("test2", srv.Client())
			// call one service
			getHelloMarket(srv,client)
			// cal other service
			serverStream(srv, client)
		case "closed":
			logger.Infof("Market Closed")
		default:
			logger.Errorf("Do not know the market status!")
		}
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
