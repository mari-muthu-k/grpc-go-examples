package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	simple_rpc "github.com/grpc-go-examples/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()
	c := simple_rpc.NewSimpleRPCClient(conn)
	//CallSimpleRPC(c)
	//CallServerStreaming(c)
	//CallClientStreaming(c)
	CallBidirectionalStreaming(c)
}

func CallSimpleRPC(c simple_rpc.SimpleRPCClient){
	// create request
	req := simple_rpc.Input{
		MyMessage: "I am running...",
	}

	// call Greet service
	res,err := c.Get(context.Background(),&req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	fmt.Println(res.MyResponse)
}

func CallServerStreaming(c simple_rpc.SimpleRPCClient){
	// create request
	req := simple_rpc.Input{
		MyMessage: "I am running...",
	}

	// call Greet service
	stream,err := c.ServerStream(context.Background(),&req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	for {
		msg,err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("received msg : ",msg);
	}
}

func CallClientStreaming(c simple_rpc.SimpleRPCClient){
	stream,err := c.ClientStream(context.Background())
	if err != nil {
		panic(err)
	}

	for i:=0;i<100;i++{
		if i%2 == 0{
			fmt.Println("request is in : ",fmt.Sprint(i))
			// create request
			req := simple_rpc.Input{
				MyMessage: "I am running in "+fmt.Sprint(i),
			}
			if err := stream.Send(&req); err != nil {
				panic(err)
			}
		}
	}

	msg,err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println("received msg : ",msg);
}

func CallBidirectionalStreaming(c simple_rpc.SimpleRPCClient){
	stream,err := c.BidirectionalStream(context.Background())
	if err != nil {
		panic(err)
	}
    recvChannel := make(chan bool)
	
	go func() {
		for {
			msg,err := stream.Recv()
			if err == io.EOF {
				recvChannel <- true
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println("Received msg : ",msg)
		}
	}()

	for i:=0;i<100;i++{
		if i%2 == 0{
			// create request
			req := simple_rpc.Input{
				MyMessage: "I am running in "+fmt.Sprint(i),
			}
			if err := stream.Send(&req); err != nil {
				panic(err)
			}
			time.Sleep(1*time.Second)
			fmt.Println("request is in : ",fmt.Sprint(i))
		}
	}

	stream.CloseSend()
	<-recvChannel
}