package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	proto.UnimplementedChittyChatServer
	messages []string
}

func (s *Server) GetMessages(ctx context.Context, in *proto.Empty) (*proto.Messages, error) {
	return &proto.Messages{Messages: s.messages}, nil
}

func main() {
	server := &Server{messages: []string{}}

	server.messages = append(server.messages, "Hello, World!")

	server.start_server()
}

func (s *Server) start_server() {
	gRPCserver := grpc.NewServer()

	netListener, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}

	proto.RegisterChittyChatServer(gRPCserver, s)

	err = gRPCserver.Serve(netListener)
	if err != nil {
		panic(err)
	}
}
