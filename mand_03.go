package main

// since it's all run locally we differ between client using their ports. If we were to run it on the internet we would use their ip addresses
import (
	"log"
	pb "."
)

// server
//serverport := 8000


func acceptConnection() {
	// accept from all
	// registrer ip:port + name + disconnect token

	//return disconnect token
}

func handleClient() {
	// handle incomming messages
	// check if they are from the right client
	// send to all other clients
}

func broadcast() {
	// send received message to all clients
}

func disconnectClient() {
	// check if ip + disconnect token received matches
	// disconnect if true
}



// client
//clientport := 9000
func setUpConnection() {
	// Send connect request to server with name + port
	// start go routine with handleClientActions
	// start go routine with receiveMessage
}

func handleClientActions() {
	// handle input from user/ random actions
	// call method of action chosen
}

func publish() {
	// send message to server
}

func receiveMessage() {
	// listen to messages
}

func disconnect() {
	// send disconnect to server
	// close client routines
}


func main() {

	server_port := 8000
	client_port := 9000

	go acceptConnection()
	log.Printf(string(server_port + client_port))
}
