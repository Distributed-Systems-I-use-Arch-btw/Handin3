package server

import (
	proto "ChittyChat/gRPC"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {

	gRPCserver := grpc.NewServer()
	netListener, err := net.Listen("tcp", ":6969")
	if err != nil {
		fmt.Println("Fool")
		panic(err)
	}
	proto.RegisterChittyChatServer(gRPCserver)
}

func bruh() {

}
