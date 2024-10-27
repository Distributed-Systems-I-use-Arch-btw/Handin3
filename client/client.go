package client

import (
	proto "ChittyChat/gRPC"
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"

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

func (c *clientInfo) updateClock(newClock *proto.LamportTimestamp) {
	if c.clock < newClock.Lamporttimestamp {
		c.clock = newClock.Lamporttimestamp
	}
}

func (c *clientInfo) GetMessage() {
	stream, _ := c.client.GetMessages(context.Background(), &proto.ClientId{Clientid: c.clientId})
	for {
		messagePackage, err := stream.Recv()
		if err != nil {
			time.Sleep(time.Millisecond)
		} else {
			log.Println(colors["green"], "Received message: ", colors["reset"], messagePackage.Message.Messages[0])
		}
	}
}

func (c *clientInfo) PostMessage(msg string) {
	c.clock += 1

	messages := &proto.Messages{Messages: []string{msg}}
	LamportTimestamp := &proto.LamportTimestamp{Lamporttimestamp: c.clock}
	postPackage := &proto.MessagePackage{Message: messages, Lamporttimestamp: LamportTimestamp}

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

func Run() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	cliInfo := &clientInfo{client: client, clientId: cliId.Clientid.Clientid, clock: cliId.Lamporttimestamp.Lamporttimestamp}

	go cliInfo.Scanner()

	cliInfo.GetMessage()

}
