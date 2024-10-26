package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChittyChatServer
	messages []string
	clock []int32
	nrClients int32
}

func (s *Server) GetMessages(ctx context.Context, in *proto.Empty) (*proto.MessagePackage, error) {
	messages := &proto.Messages{Messages: s.messages}
	vectorClock := &proto.VectorClock{Vectorclock: s.clock}
	return &proto.MessagePackage{Message: messages, Vectorclock: vectorClock}, nil
}

func (s *Server) PostMessage(ctx context.Context, in *proto.Messages) (*proto.Empty, error) {
	if len(in.Messages[0]) > 128 {
		return &proto.Empty{}, errors.New("message longer than 128 characters")
	} else if len(in.Messages[0]) == 0 {
		return &proto.Empty{}, errors.New("message is empty")
	}

	s.messages = append(s.messages, in.Messages...)
	return &proto.Empty{}, nil
}

func (s *Server) CreateClientIdentifier(ctx context.Context, in *proto.Empty) (*proto.ClientId, error) {
	s.nrClients += 1
	return &proto.ClientId{Clientid: s.nrClients}, nil
}

func main() {
	server := &Server{messages: []string{}, nrClients: 0}

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
