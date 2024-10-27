package client

import (
	proto "ChittyChat/gRPC"
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"
	"os/signal"
	"syscall"

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

func (c *clientInfo) updateClock(newTimeStamp int32) {
	if newTimeStamp > c.clock {
		c.clock = newTimeStamp
	}
	c.clock = c.clock + 1
}

func (c *clientInfo) GetMessage() {
	c.clock += 1

	stream, _ := c.client.GetMessages(context.Background(), 
		&proto.ClientPackage{
			ClientId: &proto.ClientId{Clientid: c.clientId},
			LamportTimestamp: &proto.LamportTimestamp{Lamporttimestamp: c.clock},
		})
	for {
		messagePackage, err := stream.Recv()
		if err != nil {
			time.Sleep(time.Millisecond)
		} else {
			c.updateClock(messagePackage.Lamporttimestamp.Lamporttimestamp)
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

	for true {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "exit":
			c.Exit()
		default:
			c.PostMessage(text)
		}
	}
}

var sigChan = make(chan os.Signal, 1)

func (c *clientInfo) Disconnect() {
	<-sigChan
	c.Exit()
}

func (c *clientInfo) Exit() {
	c.client.Disconnect(context.Background(), 
		&proto.ClientPackage{
			ClientId: &proto.ClientId{Clientid: c.clientId},
			LamportTimestamp: &proto.LamportTimestamp{Lamporttimestamp: c.clock},
		})
	os.Exit(0)
}

func Run() {
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChittyChatClient(conn)

	cliId, err := client.CreateClientIdentifier(context.Background(), &proto.Empty{})
	cliInfo := &clientInfo{client: client, clientId: cliId.Clientid, clock: int32(1)}

	go cliInfo.Disconnect()
	go cliInfo.Scanner()

	cliInfo.GetMessage()

}
