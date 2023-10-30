package main

//todo: comments to show understanding of wtf is going on
import (
	"bufio"
	"distributed_systems_mandatories/mandatory3/chat"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"strings"
)

var clientId int32 = -1
var timestamp int32 = 0

func main() {
	serverPort := 8000
	serverAddress := "localhost:" + strconv.Itoa(serverPort)

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer conn.Close()

	client := chat.NewChittychatClient(conn)

	stream, err := client.Connect(context.Background())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// get client id from server
	message, err := stream.Recv()
	if err != nil {
		log.Fatalf("someting wong2: %v", err)
	}
	clientId = message.ClientId
	timestamp = message.Timestamp
	log.Printf("You are now connected as client %v", message.ClientId)

	go clientListener(stream)

	for {
		reader := bufio.NewReader(os.Stdin)
		messageText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		timestamp++;
		msg := chat.Message{
			ClientId:  clientId,
			Text:      messageText,
			Timestamp: timestamp,
		}
		stream.Send(&msg)
	}
}

func clientListener(stream chat.Chittychat_ConnectClient) {
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("Client listener crashed: %v", err)
		}
		if message.Timestamp > timestamp {
			timestamp = message.Timestamp
		}

		log.Printf("%v: '%s' at time: %v", message.ClientId, strings.Replace(message.Text, "\n", "", 1), message.Timestamp)
	}
}
