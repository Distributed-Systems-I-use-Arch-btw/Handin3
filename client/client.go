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
	client   proto.ChittyChatClient
	clientId int32
	clock    []int32
}

func (c *clientInfo) max(newClock *proto.VectorClock) {
	var maxClock []int32
	var minClock []int32

	if len(c.clock) > len(newClock.GetVectorclock()) {
		maxClock = c.clock
		minClock = newClock.GetVectorclock()
	} else {
		maxClock = newClock.GetVectorclock()
		minClock = c.clock
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

	c.clock = createdClock
}

func (c *clientInfo) GetMessage() (*proto.MessagePackage, error) {
	messages, err := c.client.GetMessages(context.Background(), &proto.Empty{})
	c.clock[c.clientId] += 1
	c.max(messages.Vectorclock)
	fmt.Println(c.clock)
	return messages, err
}

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	cliInfo := &clientInfo{client: client, clientId: cliId.Clientid, clock: make([]int32, cliId.Clientid+1)}

	arg := os.Args[1]

	_, err = client.PostMessage(context.Background(), &proto.Messages{Messages: []string{arg}})
	if err != nil {
		log.Fatal(err)
	}

	messages, err := cliInfo.GetMessage()
	if err != nil {
		log.Fatal(err)
	}

	for _, messages := range messages.Message.Messages {
		fmt.Println(messages)
	}

	messages2, err := cliInfo.GetMessage()
	if err != nil {
		log.Fatal(err)
	}

	for _, messages2 := range messages2.Message.Messages {
		fmt.Println(messages2)
	}
}
