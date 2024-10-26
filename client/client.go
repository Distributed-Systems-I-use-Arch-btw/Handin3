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

func (c *clientInfo) updateClock(newClock *proto.VectorClock) {
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
	c.updateClock(messages.Vectorclock)
	fmt.Println(c.clock)
	return messages, err
}

func (c *clientInfo) PostMessage(arg string) {
	c.clock[c.clientId] += 1

	messages := &proto.Messages{Messages: []string{arg}}
	vectorClock := &proto.VectorClock{Vectorclock: c.clock}
	postPackage := &proto.MessagePackage{Message: messages, Vectorclock: vectorClock}

	c.client.PostMessage(context.Background(), postPackage)
}

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	cliInfo := &clientInfo{client: client, clientId: cliId.Clientid, clock: make([]int32, cliId.Clientid+1)}

	cliInfo.PostMessage(os.Args[1])

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
