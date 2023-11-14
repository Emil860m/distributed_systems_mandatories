package main

import (
	"distributed_systems_mandatories/mandatory4/mutex"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"slices"
	"time"
)

type mutexClient struct {
	mutex.UnimplementedMutexServer
}

var clientId = ""
var clientIp = ""
var timestamp int32 = 0
var waitingForAccess = false
var inCriticalSection = false
var replies = 0
var peerList = make([]string, 1)

// ask the other clients for access
func (client *mutexClient) askForAccess() {
	log.Printf("%v | %s's peer list is: %v", timestamp, clientId, peerList)

	log.Printf("%v | %s is now trying to gain access to the critical section\n", timestamp, clientId)

	// get the highest timestamp and set this client's timestamp to 1 higher
	for _, peer := range peerList {
		log.Printf("%v | Sending Lamport-timestamp request to %s\n", timestamp, peer)
		peerId, peerTimestamp := client.sendTimestampRequestToPeer(peer)

		log.Printf("%v | %s's timestamp was: %v", timestamp, peerId, peerTimestamp)

		if peerTimestamp > timestamp {
			timestamp = peerTimestamp
		}
	}
	timestamp++
	waitingForAccess = true
	replies = 0

	for _, peer := range peerList {
		go client.sendRequestToPeer(peer)
		log.Printf("%v | %s sent access request to %s\n", timestamp, clientId, peer)
	}

	// wait for all replies
	for replies < len(peerList) {
		time.Sleep(time.Millisecond * 100)
	}
	//replies = 0

	// Enter Critical Section
	log.Printf("%v | %s entered critical section\n", timestamp, clientId)
	inCriticalSection = true

	log.Printf("%v | %s is working in the critical section for 5 seconds", timestamp, clientId)
	time.Sleep(time.Second * 5)

	inCriticalSection = false
	waitingForAccess = false

	// Exit Critical Section
	log.Printf("%v | %s exited critical section\n", timestamp, clientId)

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
		&mutex.Request{ClientId: clientId, Timestamp: timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}

	return response.PeerList
}

func (client *mutexClient) sendTimestampRequestToPeer(peer string) (string, int32) {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	requestingClient := mutex.NewMutexClient(conn)

	response, err := requestingClient.RequestLamportTimestamp(context.Background(),
		&mutex.Request{ClientId: clientId, Timestamp: timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}
	return response.ClientId, response.Timestamp
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
		&mutex.ClientInfo{ClientId: clientId, ClientListeningIp: clientIp})
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
		ClientId:  clientId,
		Timestamp: timestamp,
	})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
		return
	}

	log.Printf("%v | '%s' granted %s access to critical section\n", timestamp, peer, clientId)

	// increment replies if we got a reply from a client ahead of this client in the queue
	replies++
}

func (client mutexClient) LetPeerKnowIExist(ctx context.Context, request *mutex.ClientInfo) (*mutex.Empty, error) {
	// add requesting client to peerList if it isn't already known
	if !slices.Contains(peerList, request.ClientListeningIp) {
		peerList = append(peerList, request.ClientListeningIp)

		log.Printf("%v | New peer added to %s's peer list: %s", timestamp, clientId, request.ClientId)
	}

	log.Printf("%v | %s let %s know they exist", timestamp, request.ClientId, clientId)
	return &mutex.Empty{}, nil
}

// RequestLamportTimestamp returns this client's Lamport timestamp
func (client mutexClient) RequestLamportTimestamp(ctx context.Context, request *mutex.Request) (*mutex.LamportTimestamp, error) {
	log.Printf("%v | %s asked for %s's timestamp: %v", timestamp, request.ClientId, clientId, timestamp)
	return &mutex.LamportTimestamp{Timestamp: timestamp}, nil
}

// RequestPeerList returns this client's list of known peerList
func (client mutexClient) RequestPeerList(ctx context.Context, request *mutex.Request) (*mutex.PeerList, error) {
	log.Printf("%v | %s sent peer list to %s: %v", timestamp, clientId, request.ClientId, peerList)
	return &mutex.PeerList{PeerList: peerList}, nil
}

// RequestAccess this is the code that is responding to other clients' requests
func (client mutexClient) RequestAccess(ctx context.Context, request *mutex.Request) (*mutex.Empty, error) {
	log.Printf("%v | %s asked %s for access to the critical section", timestamp, request.ClientId, clientId)

	// wait to respond if the requesting client is behind this client in the queue
	if inCriticalSection || (waitingForAccess && request.Timestamp > timestamp) {
		for waitingForAccess || inCriticalSection {
			time.Sleep(time.Millisecond * 100)
		}
	}

	log.Printf("%v | %s granted %s access to the critical section", timestamp, clientId, request.ClientId)

	return &mutex.Empty{}, nil
}

func setUpListener(clientIp string) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()
	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", clientIp)
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	defer listener.Close()

	log.Printf("Started listening on ip: %s\n", clientIp)

	// Register the grpc server and serve its listener
	mutex.RegisterMutexServer(grpcServer, &mutexClient{})
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run mutex_client.go <client-id> <client-ip> <known-peer-ip-1> <known-peer-ip-2> ...")
	}

	clientId = os.Args[1]
	clientIp = os.Args[2]

	if len(os.Args) > 3 {
		peerList = os.Args[3:]
	} else {
		peerList = make([]string, 0)
	}

	client := &mutexClient{}
	log.Printf("%v | Created client: %s", timestamp, clientId)

	// if there is a known peer, ask that peer for the addresses of all the other peerList
	if len(os.Args) > 3 {
		knownPeersPeerList := client.sendPeerListRequestToPeer(peerList[0])
		peerList = append(peerList, knownPeersPeerList...)

		for _, peer := range peerList {
			client.letPeerKnowIExist(peer)
		}
		log.Printf("%v | %s let peers know they exist", timestamp, clientId)
	}

	// setup listener on client port
	go setUpListener(clientIp)
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
