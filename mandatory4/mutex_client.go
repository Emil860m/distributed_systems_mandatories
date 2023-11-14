package main

import (
	"distributed_systems_mandatories/mandatory4/mutex"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
	"log"
	"net"
	"os"
	"slices"
	"time"
)

type mutexClient struct {
	clientId          string
	timestamp         int32
	waitingForAccess  bool
	inCriticalSection bool
	replies           int
	peerList          []string
	mutex.UnimplementedMutexServer
}

func newMutexClient(clientId string, peers []string) *mutexClient {
	log.Printf("0 | Created client: %s", clientId)

	return &mutexClient{
		clientId:          clientId,
		timestamp:         0,
		waitingForAccess:  false,
		inCriticalSection: false,
		replies:           0,
		peerList:          peers,
	}
}

// ask the other clients for access
func (client *mutexClient) askForAccess() {
	log.Printf("%v | %s's peer list is: %v", client.timestamp, client.clientId, client.peerList)

	// get the highest timestamp and set this client's timestamp to 1 higher
	for _, peer := range client.peerList {
		log.Printf("Sending Lamport-timestamp request to %s\n", peer)
		peerTimestamp := client.sendTimestampRequestToPeer(peer)
		log.Printf("%s's timestamp was: %v", peer, peerTimestamp)

		if peerTimestamp > client.timestamp {
			client.timestamp = peerTimestamp
		}
	}
	client.timestamp++
	log.Printf("%v | %s is now trying to gain access to the critical section\n", client.timestamp, client.clientId)

	client.waitingForAccess = true
	client.replies = 0

	for _, peer := range client.peerList {
		log.Printf("%v | Sending access request to %s\n", client.timestamp, peer)
		go client.sendRequestToPeer(peer)
	}

	//// Wait for replies
	//for i := 0; i < len(client.peerList)-1; i++ {
	//	<-time.After(time.Second) // Simulate network delay
	//}

	for client.replies < len(client.peerList) {
		time.Sleep(time.Millisecond * 100)
	}

	// Enter Critical Section
	log.Printf("%v | %s entered critical section\n", client.timestamp, client.clientId)
	client.inCriticalSection = true
	// Simulate work in the Critical Section
	log.Printf("%v | %s is working in the critical section for 5 seconds", client.timestamp, client.clientId)
	time.Sleep(time.Second * 5)
	client.inCriticalSection = false
	// Exit Critical Section
	log.Printf("%v | %s exited critical section\n", client.timestamp, client.clientId)
}

func (client *mutexClient) sendPeerListRequestToPeer(peer string) []string {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	// TODO: maybe don't create a new client? just use 'client'?
	requestingClient := mutex.NewMutexClient(conn)

	response, err := requestingClient.RequestPeerList(context.Background(),
		&mutex.Request{ClientId: client.clientId, Timestamp: client.timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}

	return response.PeerList
}

func (client *mutexClient) sendTimestampRequestToPeer(peer string) int32 {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	requestingClient := mutex.NewMutexClient(conn)

	response, err := requestingClient.RequestLamportTimestamp(context.Background(),
		&mutex.Request{ClientId: client.clientId, Timestamp: client.timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}
	return response.Timestamp
}

func (client *mutexClient) letPeerKnowIExist(peer string) {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	// TODO: maybe don't create a new client? just use 'client'?
	requestingClient := mutex.NewMutexClient(conn)

	_, err = requestingClient.LetPeerKnowIExist(context.Background(),
		&mutex.Request{ClientId: client.clientId, Timestamp: client.timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}
}

// asks for access from another specific client
func (client *mutexClient) sendRequestToPeer(peer string) {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
		return
	}
	defer conn.Close()

	requestingClient := mutex.NewMutexClient(conn)

	_, err = requestingClient.RequestAccess(context.Background(), &mutex.Request{
		ClientId:  client.clientId,
		Timestamp: client.timestamp,
	})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
		return
	}

	// increment replies if we got a reply from a client ahead of this client in the queue
	client.replies++
}

// RequestLamportTimestamp returns this client's Lamport timestamp
func (client mutexClient) LetPeerKnowIExist(ctx context.Context, request *mutex.Request) (*mutex.Empty, error) {
	// add requesting client to peerList if it is not already known
	client.addPeerToPeerList(ctx)
	log.Printf("%v | %s let %s know they exist", client.timestamp, request.ClientId, client.clientId)
	return &mutex.Empty{}, nil
}

// RequestLamportTimestamp returns this client's Lamport timestamp
func (client mutexClient) RequestLamportTimestamp(ctx context.Context, request *mutex.Request) (*mutex.LamportTimestamp, error) {
	log.Printf("%v | %s asked for %s's timestamp: %v", client.timestamp, request.ClientId, client.clientId, client.timestamp)
	return &mutex.LamportTimestamp{Timestamp: client.timestamp}, nil
}

// RequestPeerList returns this client's list of known peerList
func (client mutexClient) RequestPeerList(ctx context.Context, request *mutex.Request) (*mutex.PeerList, error) {
	log.Printf("%v | %s sent peer list to %s: %v", client.timestamp, client.clientId, request.ClientId, client.peerList)
	return &mutex.PeerList{PeerList: client.peerList}, nil
}

// RequestAccess this is the code that is responding to other clients' requests
func (client mutexClient) RequestAccess(ctx context.Context, request *mutex.Request) (*mutex.Empty, error) {
	log.Printf("%v | %s asked %s for access to the critical section", client.timestamp, request.ClientId, client.clientId)

	// add requesting client to peerList if it is not already known
	client.addPeerToPeerList(ctx)

	// wait to respond if the requesting client is behind this client in the queue
	if client.inCriticalSection || client.waitingForAccess && request.Timestamp < client.timestamp {
		for client.waitingForAccess == true {
			time.Sleep(time.Millisecond * 100)
		}
	}

	log.Printf("%v | %s granted %s access to the critical section", client.timestamp, client.clientId, request.ClientId)

	return &mutex.Empty{}, nil
}

func (client mutexClient) addPeerToPeerList(ctx context.Context) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatal("Failed to get requesting peer's address")
	}
	p.Addr
	clientIp := p.Addr.String()

	if !slices.Contains(client.peerList, clientIp) {
		client.peerList = append(client.peerList, clientIp)

		log.Printf("New peer IP added to list: %s", clientIp)
	}
}

func setUpListener(clientPort string) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()
	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", clientPort))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	defer listener.Close()

	log.Printf("Started listening on port: %s\n", clientPort)

	// Register the grpc server and serve its listener
	mutex.RegisterMutexServer(grpcServer, &mutexClient{})
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatal("Usage: go run mutex_client.go <client-id> <client-port> <known-peer-ip>")
	}

	clientId := os.Args[1]
	clientPort := os.Args[2]
	peerList := make([]string, 0)

	client := newMutexClient(clientId, peerList)

	// if there is a known peer, ask that peer for the addresses of all the other peerList
	if len(os.Args) == 4 {
		knownPeer := os.Args[3]
		client.peerList = append(client.peerList, knownPeer)

		knownPeersPeerList := client.sendPeerListRequestToPeer(knownPeer)
		client.peerList = append(client.peerList, knownPeersPeerList...)

		for _, peer := range client.peerList {
			client.letPeerKnowIExist(peer)
		}
		log.Printf("%v | %s let peers know they exist", client.timestamp, client.clientId)

	}

	// setup listener on client port
	go setUpListener(clientPort)
	// sleep so listener is completely set up
	time.Sleep(time.Second)

	for {
		fmt.Println("\n--- Press [ENTER] to ask for access to the critical section ---")
		_, err := fmt.Scanln()
		if err != nil {
			log.Fatal(err)
			return
		}
		client.askForAccess()
	}
}
