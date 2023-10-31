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

var clientName = ""
var clientTimestamp int32 = 0

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
		log.Fatalf("Counldn't connect to server: %v", err)
	}
	clientName = message.ClientName
	clientTimestamp = message.Timestamp
	log.Printf("You are now connected as %v", message.ClientName)

	go clientListener(stream)

	for {
		reader := bufio.NewReader(os.Stdin)
		messageText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		// Normally the server should also check if the message is to long
		if len(messageText) > 128 {
			log.Printf("Error: You message is to long. Limit: 128 characters")
			continue
		}
		clientTimestamp++
		msg := chat.Message{
			ClientName: clientName,
			Text:       strings.Replace(messageText, "\n", "", 1),
			Timestamp:  clientTimestamp,
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
		if message.Timestamp > clientTimestamp {
			clientTimestamp = message.Timestamp
		}

		log.Printf("%v: '%s' (time: %v)", message.ClientName, message.Text, message.Timestamp)
	}
}
