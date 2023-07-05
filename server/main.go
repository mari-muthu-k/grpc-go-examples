package main

import (
	"context"
	"io"
	"log"
	"net"

	simple_rpc "github.com/grpc-go-examples/proto"
	"google.golang.org/grpc"
)

type Server struct{}

func (*Server)Get(ctx context.Context,req *simple_rpc.Input)(*simple_rpc.Output,error){
	myMessage := req.GetMyMessage()
	res := simple_rpc.Output{
		MyResponse: "okay, so your message is -> "+myMessage,
	}
	return &res,nil
}

func (*Server)ServerStream(req *simple_rpc.Input,stream simple_rpc.SimpleRPC_ServerStreamServer)error{
	myMessage := req.GetMyMessage()
	res := simple_rpc.Output{
		MyResponse: "okay, so your message is in even number -> "+myMessage,
	}
	
	for i:=0;i<100;i++{
		if i%2 == 0 {
			stream.Send(&res)
		}
	}
	return nil
}

func (*Server)ClientStream(reqStream simple_rpc.SimpleRPC_ClientStreamServer)error{
	var messages []string
	for {
		req,err := reqStream.Recv()
		if err == io.EOF {
			return reqStream.SendAndClose(&simple_rpc.MessageList{Output: messages})
		}
		if err != nil {
			panic(err)
		}
		myMessage := req.GetMyMessage()
		res := "okay, so your message is -> "+myMessage
		messages = append(messages, res)
	}
}

func (*Server)BidirectionalStream(reqStream simple_rpc.SimpleRPC_BidirectionalStreamServer)error{
	for {
		req,err := reqStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			panic(err)
		}
		myMessage := req.GetMyMessage()
		res := simple_rpc.Output{
			MyResponse: "okay, so your message is -> "+myMessage,
		}
		if err = reqStream.Send(&res); err != nil {
			log.Fatalf("err in sending : ",err)
		}
	}
}

func main(){
	net,err := net.Listen("tcp","localhost:8080"); if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	simple_rpc.RegisterSimpleRPCServer(s,&Server{})

	if err := s.Serve(net); err != nil {
		panic(err)
	}
	}