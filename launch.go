package main
	
import (
    "os"
	"fmt"
	client "ChittyChat/client"
	server "ChittyChat/server"
)

func main() {
	arg := os.Args

	fmt.Println(arg[1])

	switch arg[1] {
	case "client":
		client.Run()
	case "server":
		server.Run()
	}
}