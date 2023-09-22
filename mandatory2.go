package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverChannel := make(chan bool)
	clientChannel := make(chan bool)

	go server(serverChannel)
	go client(clientChannel)

	<-serverChannel
	<-clientChannel
}
func server(completeChannel chan bool) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on 127.0.0.1:8080")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connection established with client")

	completeChannel <- true
	// Handle incoming data or perform actions here
}

func client(completeChannel chan bool) {
	time.Sleep(time.Second * 4)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	completeChannel <- true
	// Perform any necessary actions here or send data to the server
}
