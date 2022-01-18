package main

import (
	"context"
	// "encoding/json"
	"fmt"
	// "time"
	// "strconv"

	pb "github.com/egargale/tradier_gomicro_backend/services/test2/proto"

	"github.com/micro/micro/v3/service"
	//"github.com/micro/micro/v3/service/broker"

)

type Healthcheck struct {
	Healthy bool
	Service string
	Notification string
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

	// create the proto client for helloworld
	client := pb.NewTest2Service("test2", srv.Client())

	// call an endpoint on the service

	rsp, err := client.Stream(context.Background(), &pb.StreamingRequest{
		Count: 200,
	})
	if err != nil {
		fmt.Println("Error calling mytest: ", err)
		return
	}
	

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
	fmt.Println("Response: ", rsp)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)

}
