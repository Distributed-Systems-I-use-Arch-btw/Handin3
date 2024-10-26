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
	clientId int
}

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	fmt.Println(cliId)
	// cliInfo := &clientInfo{client: cliId}
	// _ = cliInfo

	arg := os.Args[1]

	_, err = client.PostMessage(context.Background(), &proto.Messages{Messages: []string{arg}})
	if err != nil {
		log.Fatal(err)
	}

	messages, err := client.GetMessages(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for _, messages := range messages.Message.Messages {
		fmt.Println(messages)
	}
}
