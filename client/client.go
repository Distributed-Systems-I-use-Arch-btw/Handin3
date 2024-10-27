package main

import (
	proto "ChittyChat/gRPC"
	"context"
	"fmt"
	"log"
	"time"
	"bufio"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientInfo struct {
	client   proto.ChittyChatClient
	clientId int32
	clock    []int32
}

var colors = map[string]string{
    "red":    "\033[31m",
    "green":  "\033[32m",
    "yellow": "\033[33m",
    "blue":   "\033[34m",
    "reset":  "\033[0m",
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
	stream, _ := c.client.GetMessages(context.Background(), &proto.ClientId{Clientid: c.clientId})
	for {
		messagePackage, err := stream.Recv()
		if err != nil {
            time.Sleep(time.Second)
        } else {
			fmt.Println(messagePackage.Vectorclock.Vectorclock)
			fmt.Println(colors["green"], "Received message: ", colors["reset"], messagePackage.Message.Messages)
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
	reader := bufio.NewReader(os.Stdin)
	running := true

	for running {
		
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
			case "exit":
				running = false
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


