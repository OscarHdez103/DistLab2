package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "3.91.227.58:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	request := stubs.Request{Message: "Hello World"}
	response := new(stubs.Response)
	client.Call(stubs.ReverseHandler, request, response)
	fmt.Println("Responded: " + response.Message)

}
