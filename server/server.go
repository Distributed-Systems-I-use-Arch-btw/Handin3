package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"errors"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChittyChatServer
	messages  []string
	clock     []int32
	nrClients int32
}

func (s *Server) updateClock(newClock *proto.VectorClock) {
	var maxClock []int32
	var minClock []int32

	if len(s.clock) > len(newClock.GetVectorclock()) {
		maxClock = s.clock
		minClock = newClock.GetVectorclock()
	} else {
		maxClock = newClock.GetVectorclock()
		minClock = s.clock
	}

	createdClock := make([]int32, len(maxClock))

	for i := 0; i < len(minClock); i++ {
		if maxClock[i] > minClock[i] {
			createdClock[i] = maxClock[i]
		} else {
			createdClock[i] = minClock[i]
		}
	}

	for i := len(minClock); i < len(maxClock); i++ {
		createdClock[i] = maxClock[i]
	}

	s.clock = createdClock
}

func (s *Server) GetNewMessages(oldMessages []string) (NewMessages []string) {
	if len(oldMessages) < len(s.messages) {
		return s.messages[len(oldMessages):]
	}
	return []string{}
}

func (s *Server) GetMessages(in *proto.Empty, stream proto.ChittyChat_GetMessagesServer) error {
	s.clock[0] += 1
	oldMessages := s.messages
	for _, message := range oldMessages {
        messagePackage := &proto.MessagePackage{
            Message:     &proto.Messages{Messages: []string{message}},
            Vectorclock: &proto.VectorClock{Vectorclock: s.clock},
        }

        if err := stream.Send(messagePackage); err != nil {
            return err
        }
    }
	for {
		time.Sleep(time.Second)

		newMessage := s.GetNewMessages(oldMessages)
		oldMessages = s.messages
		
		for _, message := range newMessage {
			messagePackage := &proto.MessagePackage{
				Message:     &proto.Messages{Messages: []string{message}},
				Vectorclock: &proto.VectorClock{Vectorclock: s.clock},
			}
	
			if err := stream.Send(messagePackage); err != nil {
				return err
			}
		}
		select {
			case <-stream.Context().Done():
				return nil
			default:
				continue
        }
	}
	// messages := &proto.Messages{Messages: s.messages}
	// vectorClock := &proto.VectorClock{Vectorclock: s.clock}
}

func (s *Server) PostMessage(ctx context.Context, in *proto.MessagePackage) (*proto.Empty, error) {
	s.clock[0] += 1

	curMessage := in.Message.Messages

	if len(curMessage[0]) > 128 {
		return &proto.Empty{}, errors.New("message longer than 128 characters")
	} else if len(curMessage[0]) == 0 {
		return &proto.Empty{}, errors.New("message is empty")
	}

	s.messages = append(s.messages, curMessage[0])

	s.updateClock(in.GetVectorclock())

	return &proto.Empty{}, nil
}

func (s *Server) CreateClientIdentifier(ctx context.Context, in *proto.Empty) (*proto.ClientId, error) {
	s.nrClients += 1
	return &proto.ClientId{Clientid: s.nrClients}, nil
}

func main() {
	server := &Server{messages: []string{}, clock: make([]int32, 5), nrClients: 0}

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
