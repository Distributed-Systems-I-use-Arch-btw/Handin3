package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	messages, err := client.GetMessages(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for _, messages := range messages.Messages {
		fmt.Println(messages)
	}
}
