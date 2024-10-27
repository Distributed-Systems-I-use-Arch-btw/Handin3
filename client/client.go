package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"fmt"
	"log"
	"time"

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

func (c *clientInfo) GetMessage() {
	stream, _ := c.client.GetMessages(context.Background(), &proto.Empty{})
	for {
		messagePackage, err := stream.Recv()
		if err != nil {
            time.Sleep(time.Second)
        } else {
			fmt.Println(messagePackage.Vectorclock.Vectorclock)
			fmt.Println("Received message: ", messagePackage.Message.Messages)
		}
	}
}

func (c *clientInfo) PostMessage(msg string) {
	c.clock[c.clientId] += 1

	messages := &proto.Messages{Messages: []string{msg}}
	vectorClock := &proto.VectorClock{Vectorclock: c.clock}
	postPackage := &proto.MessagePackage{Message: messages, Vectorclock: vectorClock}

	c.client.PostMessage(context.Background(), postPackage)
}

func (c *clientInfo) Scanner() {
	running := true
	var text string
	for running {
		fmt.Scan(&text)
		switch text {
			case "exit":
				break
			default:
				c.PostMessage(text)
		}
	}
}

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	cliInfo := &clientInfo{client: client, clientId: cliId.Clientid, clock: make([]int32, cliId.Clientid+1)}

	go cliInfo.Scanner()
	
	cliInfo.GetMessage()

}


