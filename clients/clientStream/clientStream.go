package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	pb "github.com/egargale/tradier_gomicro_backend/services/test2/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/broker"
)

type Healthcheck struct {
	Healthy      bool
	Service      string
	Notification string
}

// func bidirectional(cl proto.StreamerService) {

// // create streaming client
// 	stream, err := cl.Stream(context.Background())
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}

// 	// bidirectional stream
// 	// send and receive messages for a 10 count
// 	for j := 0; j < 10; j++ {
// 		if err := stream.Send(&proto.Request{Count: int64(j)}); err != nil {
// 			fmt.Println("err:", err)
// 			return
// 		}
// 		rsp, err := stream.Recv()
// 		if err != nil {
// 			fmt.Println("recv err", err)
// 			break
// 		}
// 		fmt.Printf("Sent msg %v got msg %v\n", j, rsp.Count)
// 	}

// 	// close the stream
// 	if err := stream.Close(); err != nil {
// 		fmt.Println("stream close err:", err)
// 	}
// }

func publishStream(resp pb.StreamingResponse) error {

	bytes, err := json.Marshal(&Healthcheck{
		Healthy:      true,
		Service:      "foo",
		Notification: strconv.Itoa(int(resp.Count)),
	})

	if err != nil {
		fmt.Println("Error publishing ", err)
		return err
	}

	stream_err := broker.Publish("health", &broker.Message{Body: bytes})
	if err != nil {
		fmt.Println("Message not brocasted: Error", stream_err)
		return stream_err
	}
	return nil
}

func serverStream(cl pb.Test2Service) {
	// send request to stream count of 10
	stream, err := cl.Stream(context.Background(), &pb.StreamingRequest{Count: int64(10)})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var i int

	// server side stream
	// receive messages for a 10 count
	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			break
		}
		i++
		fmt.Printf("got msg %v\n", rsp.Count)

		// Publish to broker
		pub_err := publishStream(*rsp)
		if pub_err != nil {
			fmt.Println("Error in broadcasting: Error", pub_err)
		}
	}

	if i < 10 {
		fmt.Println("only got", i)
		return
	}

	// close the stream
	if err := stream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
}

func main() {
	// create and initialise a new service
	srv := service.New()

	// Load the config
	// val, err := config.Get("tradier.production.account")
	// if err != nil {
	// 	logger.Fatalf("Could not load mykey: %v", err)
	// } else {
	// 	logger.Infof("mykey = %v", val)
	// }

	// create the proto client for
	client := pb.NewTest2Service("test2", srv.Client())

	// call an endpoint on the service

	// rsp, err := client.Stream(context.Background(), &pb.StreamingRequest{
	// 	Count: 200,
	// })
	// if err != nil {
	// 	fmt.Println("Error calling streaming server: ", err)
	// 	return
	// }

	// bytes, err := json.Marshal(&Healthcheck{
	// 	Healthy: true,
	// 	Service: "foo",
	// 	Notification: strconv.Itoa(rsp.RecvMsg()),
	// })

	// if err != nil {
	// 	fmt.Println("Error publishing ", err)
	// 	return
	// }

	// broker.Publish("health", &broker.Message{Body: bytes})

	// print the response
	// fmt.Println("Response: ", rsp)

	// let's delay the process for exiting for reasons you'll see below
	// time.Sleep(time.Second * 5)

	for {
		// fmt.Println("Stream")
		// // bidirectional stream
		// bidirectional(cl)

		fmt.Println("ServerStream")
		// server side stream
		serverStream(client)

		time.Sleep(time.Second)
	}
}
