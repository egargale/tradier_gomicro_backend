package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	proto "github.com/micro/services/helloworld/proto"
)

func call(i int) {
	// create and initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("helloworld", srv.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &proto.CallRequest{
		Name: "Enrico",
	})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}

	// print the response
	fmt.Printf("Count %d Response: %s \n", i, rsp.Message)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}

// func stream(i int) {
// 	// Create new request to service go.micro.srv.example, method Example.Call
// 	// Request can be empty as its actually ignored and merely used to call the handler
// 	req := client.NewRequest("tcore", "Tcore.Stream", &tcore.StreamingRequest{})

// 	stream, err := client.Stream(context.Background(), req)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	if err := stream.Send(&tcore.StreamingRequest{Count: int64(i)}); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	for stream.Error() == nil {
// 		rsp := &tcore.StreamingResponse{}
// 		err := stream.Recv(rsp)
// 		if err != nil {
// 			fmt.Println("recv err", err)
// 			break
// 		}
// 		fmt.Println("Stream: rsp:", rsp.Count)
// 	}

// 	if stream.Error() != nil {
// 		fmt.Println("stream err:", err)
// 		return
// 	}

// 	if err := stream.Close(); err != nil {
// 		fmt.Println("stream close err:", err)
// 	}
// }

func main() {

	for i := 0; i < 10; i++ {
		call(i)
	}
}
