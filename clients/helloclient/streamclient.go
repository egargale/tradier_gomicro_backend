package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	proto "github.com/micro/services/helloworld/proto"
)

func stream(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	// Request can be empty as its actually ignored and merely used to call the handler
	req := client.NewRequest("go.micro.srv.example", "Example.Stream", &example.StreamingRequest{})

	stream, err := client.Stream(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if err := stream.Send(&example.StreamingRequest{Count: int64(i)}); err != nil {
		fmt.Println("err:", err)
		return
	}
	for stream.Error() == nil {
		rsp := &example.StreamingResponse{}
		err := stream.Recv(rsp)
		if err != nil {
			fmt.Println("recv err", err)
			break
		}
		fmt.Println("Stream: rsp:", rsp.Count)
	}

	if stream.Error() != nil {
		fmt.Println("stream err:", err)
		return
	}

	if err := stream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
}


func main() {
	// create and initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("helloworld", srv.Client())

	// call an endpoint on the service
	rsp, err := client.Stream(context.Background(), &proto.StreamRequest{
		Name: "John",
	})
	
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.RecvMsg())

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}
