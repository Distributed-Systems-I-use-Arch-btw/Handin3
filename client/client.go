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
	clock    int32
}

var colors = map[string]string{
    "red":    "\033[31m",
    "green":  "\033[32m",
    "yellow": "\033[33m",
    "blue":   "\033[34m",
    "reset":  "\033[0m",
} 

func (c *clientInfo) updateClock(newClock *proto.VectorClock) {
	if c.clock < newClock.Vectorclock {
		c.clock = newClock.Vectorclock
	}
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
	c.clock += 1

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
	cliInfo := &clientInfo{client: client, clientId: cliId.Clientid, clock: int32(0)}

	go cliInfo.Scanner()
	
	cliInfo.GetMessage()

}


