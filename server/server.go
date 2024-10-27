package server

import (
	proto "ChittyChat/gRPC"
	"context"
	"errors"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChittyChatServer
	nrClients int32
	msData    timedMessages
}
type timedMessages struct {
	messages   []string
	timeStamps []int32
}

func (s *Server) GetNewMessages(oldMessagesLen int) (NewMessages *timedMessages) {
	if oldMessagesLen < len(s.msData.messages) {
		return &timedMessages{
			messages:   s.msData.messages[oldMessagesLen:],
			timeStamps: s.msData.timeStamps[oldMessagesLen:],
		}
	}
	return &timedMessages{
		messages:   []string{},
		timeStamps: []int32{},
	}
}

func streamMessages(sendMessages timedMessages, stream proto.ChittyChat_GetMessagesServer, s *Server) {
	for i := 0; i < len(sendMessages.messages); i++ {
		messagePackage := &proto.MessagePackage{
			Message:          &proto.Messages{Messages: []string{sendMessages.messages[i]}},
			Lamporttimestamp: &proto.LamportTimestamp{Lamporttimestamp: sendMessages.timeStamps[i]},
		}

		if err := stream.Send(messagePackage); err != nil {
			break
		}
	}
}

func (s *Server) GetMessages(clientInfo *proto.ClientPackage, stream proto.ChittyChat_GetMessagesServer) error {
	currentMessages := &s.msData
	length := len(currentMessages.messages)
	streamMessages(*currentMessages, stream, s)

	log.Println("Received GetMessages call from Client Id " + strconv.Itoa(int(clientInfo.LamportTimestamp.Lamporttimestamp)) + " at Lamport time " +  "TBA")

	for {
		time.Sleep(time.Millisecond)

		currentMessages = s.GetNewMessages(length)
		length = len(s.msData.messages)

		streamMessages(*currentMessages, stream, s)

		select {
		case <-stream.Context().Done():
			hasLeft := "Participant " + strconv.Itoa(int(clientInfo.ClientId.Clientid)) + " left Chitty-Chat at Lamport time " + "TBA"
			log.Println(hasLeft)
			s.msData.messages = append(s.msData.messages, hasLeft)
			s.msData.timeStamps = append(s.msData.timeStamps, int32(0))
			return nil
		default:
			continue
		}
	}
}

func (s *Server) PostMessage(ctx context.Context, in *proto.MessagePackage) (*proto.Empty, error) {
	curMessage := in.Message.Messages

	if len(curMessage[0]) > 128 {
		return &proto.Empty{}, errors.New("message longer than 128 characters")
	} else if len(curMessage[0]) == 0 {
		return &proto.Empty{}, errors.New("message is empty")
	}

	s.msData.messages = append(s.msData.messages, curMessage[0])
	s.msData.timeStamps = append(s.msData.timeStamps, int32(0))

	log.Println("Received PostMessage call at Lamport time " + strconv.Itoa(int(in.Lamporttimestamp.Lamporttimestamp)))

	return &proto.Empty{}, nil
}

func (s *Server) CreateClientIdentifier(ctx context.Context, in *proto.Empty) (*proto.ClientId, error) {
	s.nrClients += 1

	hasJoined := "Participant " + strconv.Itoa(int(s.nrClients)) + " joined Chitty-Chat at Lamport time " + "TBA"
	log.Println(hasJoined)

	s.msData.messages = append(s.msData.messages, hasJoined)
	s.msData.timeStamps = append(s.msData.timeStamps, int32(0))
	return &proto.ClientId{Clientid: s.nrClients}, nil
}

func Run() {
	server := &Server{
		nrClients: 0,
		msData: timedMessages{
			messages:   make([]string, 0),
			timeStamps: make([]int32, 0),
		},
	}

	server.start_server()
}

func (s *Server) start_server() {

	gRPCserver := grpc.NewServer()

	log.Println("Server started")

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
