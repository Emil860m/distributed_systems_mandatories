package main

//todo: A valid message is a string of UTF-8 encoded text with a maximum length of 128 characters
//todo: comments to show understanding of wtf is going on
import (
	"distributed_systems_mandatories/mandatory3/chat"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

var clientStreams []chat.Chittychat_ConnectServer
var highestClientId int32 = 0
var serverTimestamp int32 = 0
var serverPort int = 8000

func main() {

	server := &ChittychatServer{}

	startServer(server)

	time.Sleep(time.Hour * 10)
}

// we used this for help
// https://github.com/Mai-Sigurd/grpcTimeRequestExample?tab=readme-ov-file#setting-up-the-server
func startServer(server *ChittychatServer) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(serverPort))
	defer listener.Close()

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", serverPort)

	// Register the grpc server and serve its listener
	chat.RegisterChittychatServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

type ChittychatServer struct {
	chat.UnimplementedChittychatServer
}

func (s ChittychatServer) Connect(stream chat.Chittychat_ConnectServer) error {
	highestClientId++
	clientName := "Client " + strconv.Itoa(int(highestClientId))
	log.Printf("Client %v has connected", highestClientId)

	clientStreams = append(clientStreams, stream)
	serverTimestamp++

	stream.Send(&chat.Message{
		ClientName: clientName,
		Text:       "You are now connected",
		Timestamp:  serverTimestamp,
	})

	go broadcastMessage(chat.Message{
		ClientName: "Server",
		Text:       clientName + " has connected",
		Timestamp:  serverTimestamp,
	})

	serverListener(stream, clientName)

	return nil
}

func serverListener(stream chat.Chittychat_ConnectServer, clientName string) {
	for {
		message, err := stream.Recv()

		if err != nil {
			clientStreams = removeStreamFromList(clientStreams, stream)

			log.Printf("Client %v had disconnected", clientName)
			serverTimestamp++
			broadcastMessage(chat.Message{
				ClientName: clientName,
				Text:       clientName + " had disconnected",
				Timestamp:  serverTimestamp,
			})

			return
		}

		if message.Timestamp > serverTimestamp {
			serverTimestamp = message.Timestamp
		}
		go broadcastMessage(*message)
	}
}

func removeStreamFromList(clientList []chat.Chittychat_ConnectServer, streamToRemove chat.Chittychat_ConnectServer) []chat.Chittychat_ConnectServer {
	for i := 0; i < len(clientList); i++ {
		if clientList[i] == streamToRemove {
			clientList[i] = clientList[len(clientList)-1]
			return clientList[:len(clientList)-1]
			//break
		}
	}
	return clientList
}

func broadcastMessage(message chat.Message) {
	log.Printf("Broadcasting message: '%v' from %v (time: %v)", message.Text, message.ClientName, serverTimestamp)
	for i := 0; i < len(clientStreams); i++ {
		clientStreams[i].Send(&message)
	}
}
