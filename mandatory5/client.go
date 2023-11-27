package main

import (
	"bufio"
	"distributed_systems_mandatories/mandatory5/auction"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"strings"
)

var username string
var serverIP string

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run client.go <username> <server IP>")
	}

	username = os.Args[1]
	serverIP = os.Args[2]
	fmt.Println("Usage\n- bid <amount>\n- result")

	for {

		text := takeInput("Enter your command: ")
		args := strings.Split(text, " ")
		if len(args) == 2 && args[0] == "bid" {
			makeBid(args[1])
		} else if len(args) == 1 && args[0] == "result" {
			getResult()
		} else {
			fmt.Println("Usage\n- bid <amount>\n- result")
		}
	}
}

func takeInput(inputCommand string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\n%v", inputCommand)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}
	return strings.TrimSpace(strings.ToLower(text))
}

func makeBid(amountStr string) {
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		log.Fatalf("%v is not a number!", amount)
	}
	amount32 := int32(amount)
	fmt.Printf("Making bid %v as %s\n", amount32, username)

	conn, err := grpc.Dial(serverIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error connecting to server %s: %v", serverIP, err)
		text := takeInput("Enter new server IP and port: ")
		serverIP = text
		makeBid(amountStr)
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)
	response, err := requestingClient.Bid(context.Background(), &auction.BidMessage{
		Name:   username,
		Amount: amount32,
	})
	if err != nil {
		log.Printf("Error connecting to server %s: %v", serverIP, err)
		text := takeInput("Enter new server IP and port: ")
		serverIP = text
		makeBid(amountStr)
		return
	}
	fmt.Printf("Response: %v", response)
}
func getResult() {
	conn, err := grpc.Dial(serverIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error connecting to server %s: %v", serverIP, err)
		text := takeInput("Enter new server IP and port: ")
		serverIP = text
		getResult()
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)
	response, err := requestingClient.Result(context.Background(), &auction.Empty{})
	if err != nil {
		log.Printf("Error connecting to server %s: %v", serverIP, err)
		text := takeInput("Enter new server IP and port: ")
		serverIP = text
		getResult()
		return
	}
	if response.Ongoing {
		fmt.Println("Auction still ongoing")
	} else {
		fmt.Println("Auction finished")
	}
	fmt.Printf("Highest bidder: %v\n", response.Name)
	fmt.Printf("Highest bid: %v\n", response.Amount)
}
