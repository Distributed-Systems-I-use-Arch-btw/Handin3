package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"errors"
	"net"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChittyChatServer
	clock     int32
	nrClients int32
	msData	  timedMessages
}
type timedMessages struct {
	messages	[]string
	timeStamps	[]int32
}

func (s *Server) updateClock(newClock *proto.VectorClock) {
	if s.clock < newClock.Vectorclock {
		s.clock = newClock.Vectorclock
	}
}

func (s *Server) GetNewMessages(oldMessagesLen int) (NewMessages *timedMessages) {
	if oldMessagesLen < len(s.msData.messages) {
		return &timedMessages{
			messages: s.msData.messages[oldMessagesLen:],
			timeStamps: s.msData.timeStamps[oldMessagesLen:],
		}
	}
	return &timedMessages{
		messages: []string{},
		timeStamps: []int32{},
	}
}

func streamMessages(sendMessages timedMessages, stream proto.ChittyChat_GetMessagesServer, s *Server) {
	for i := 0; i < len(sendMessages.messages); i++ {
        messagePackage := &proto.MessagePackage{
            Message:     &proto.Messages{Messages: []string{sendMessages.messages[i]}},
            Vectorclock: &proto.VectorClock{Vectorclock: sendMessages.timeStamps[i]},
        }

        if err := stream.Send(messagePackage); err != nil {
            break
        }
    }
}

func (s *Server) GetMessages(id *proto.ClientId, stream proto.ChittyChat_GetMessagesServer) error {
	s.clock += 1
	currentMessages := &s.msData
	length := len(currentMessages.messages)
	streamMessages(*currentMessages, stream, s)

	for {
		time.Sleep(time.Second)

		currentMessages = s.GetNewMessages(length)
		length = len(s.msData.messages)
		
		streamMessages(*currentMessages, stream, s)

		select {
			case <-stream.Context().Done():
				hasLeft := fmt.Sprintf("Participant %d left Chitty-Chat at Vector time z", id.Clientid) 
				//Might need to update vector clock?
				s.msData.messages = append(s.msData.messages, hasLeft)
				s.msData.timeStamps = append(s.msData.timeStamps, s.clock)
				return nil
			default:
				continue
        }
	}
}

func (s *Server) PostMessage(ctx context.Context, in *proto.MessagePackage) (*proto.Empty, error) {
	s.clock += 1

	curMessage := in.Message.Messages

	if len(curMessage[0]) > 128 {
		return &proto.Empty{}, errors.New("message longer than 128 characters")
	} else if len(curMessage[0]) == 0 {
		return &proto.Empty{}, errors.New("message is empty")
	}

	s.updateClock(in.GetVectorclock())

	s.msData.messages = append(s.msData.messages, curMessage[0])
	s.msData.timeStamps = append(s.msData.timeStamps, s.clock)

	return &proto.Empty{}, nil
}

func (s *Server) CreateClientIdentifier(ctx context.Context, in *proto.Empty) (*proto.ClientId, error) {
	s.nrClients += 1
	//Might need to update vector clock?
	hasJoined := fmt.Sprintf("Participant %d joined Chitty-Chat at Vector time z", s.nrClients) 
	
	s.msData.messages = append(s.msData.messages, hasJoined)
	s.msData.timeStamps = append(s.msData.timeStamps, s.clock)
	return &proto.ClientId{Clientid: s.nrClients}, nil
}

func main() {
	server := &Server{
		clock: int32(0), 
		nrClients: 0,
		msData: timedMessages{
			messages: make([]string, 0),
			timeStamps: make([]int32, 0),
		}, 
	}

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
