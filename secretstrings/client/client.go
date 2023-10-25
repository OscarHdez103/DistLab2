package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "3.88.159.164:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, err := os.Open("wordlist")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()
		request := stubs.Request{Message: line}
		response := new(stubs.Response)
		client.Call(stubs.ReverseHandler, request, response)
		fmt.Println("Responded: " + response.Message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}
