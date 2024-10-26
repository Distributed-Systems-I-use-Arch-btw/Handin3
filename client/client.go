package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientInfo struct {
	clientId int32
	clock []int32
}

func (c *clientInfo) GetMessage (ctx context.Context, client proto.ChittyChatClient) (*proto.MessagePackage, error) {
	messages, err := client.GetMessages(context.Background(), &proto.Empty{})
	// max(c.clock, messages.Vectorclock)
	fmt.Println(messages.Vectorclock)
	return messages, err
}

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)
	_ = client

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	cliInfo := &clientInfo{clientId: cliId.Clientid}
	fmt.Println(cliInfo)

	arg := os.Args[1]

	_, err = client.PostMessage(context.Background(), &proto.Messages{Messages: []string{arg}})
	if err != nil {
		log.Fatal(err)
	}

	messages, err := cliInfo.GetMessage(context.Background(), client)

	for _, messages := range messages.Message.Messages {
		fmt.Println(messages)
	}
}
