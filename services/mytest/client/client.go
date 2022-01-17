package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	proto "github.com/egargale/tradier_gomicro_backend/services/mytest/proto"
)

func main() {
	// create and initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewMytestService("mytest", srv.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println("Error calling mytest: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.Response)
	
	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}
