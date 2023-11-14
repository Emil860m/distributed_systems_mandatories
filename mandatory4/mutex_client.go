package main

import (
	"distributed_systems_mandatories/mandatory4/mutex"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

type MutexClient struct {
	clientId          string
	timestamp         int64
	waitingForAccess  bool
	inCriticalSection bool
	replies           int
	peers             []string
}

func NewMutexClient(clientId string, peers []string) *MutexClient {
	return &MutexClient{
		clientId:          clientId,
		timestamp:         0,
		waitingForAccess:  false,
		inCriticalSection: false,
		replies:           0,
		peers:             peers,
	}
}

// ask the other clients for access
func (mc *MutexClient) askForAccess() {
	log.Printf("%s is now trying to get access\n", mc.clientId)

	mc.waitingForAccess = true
	mc.timestamp++
	mc.replies = 0

	for _, peer := range mc.peers {
		log.Printf("Sending access request to '%s'\n", peer)
		go mc.sendRequestToPeer(peer)
	}

	//// Wait for replies
	//for i := 0; i < len(mc.peers)-1; i++ {
	//	<-time.After(time.Second) // Simulate network delay
	//}

	for mc.replies < len(mc.peers) {
		time.Sleep(time.Millisecond * 100)
	}

	// Enter Critical Section
	fmt.Printf("Node %s entered Critical Section\n", mc.clientId)
	mc.inCriticalSection = true
	// Simulate work in the Critical Section
	time.Sleep(time.Second)
	mc.inCriticalSection = false
	// Exit Critical Section
	fmt.Printf("Node %s exited Critical Section\n", mc.clientId)

	//// Release access to other nodes
	//for _, peer := range mc.peers {
	//	go mc.sendReleaseToPeer(peer)
	//}
}

// asks for access from another specific client
func (mc *MutexClient) sendRequestToPeer(peer string) {
	conn, err := grpc.Dial(peer)
	if err != nil {
		log.Printf("Error connecting to peer %s: %v", peer, err)
		return
	}
	defer conn.Close()

	// TODO: maybe don't create a new client? just use 'mc'?
	client := mutex.NewMutexClient(conn)

	response, err := client.RequestAccess(context.Background(), &mutex.Request{
		ClientId:  mc.clientId,
		Timestamp: mc.timestamp,
	})
	if err != nil {
		log.Printf("Error requesting access from peer %s: %v", peer, err)
		return
	}

	if response.Granted {
		mc.replies++
	}
}

// RequestAccess this is the code that is responding to other clients' requests
func (mc *MutexClient) RequestAccess(ctx context.Context, request *mutex.Request) (*mutex.Response, error) {
	if mc.inCriticalSection || mc.waitingForAccess && request.Timestamp < mc.timestamp {
		for mc.waitingForAccess == true {
			time.Sleep(time.Millisecond * 100)
		}
	}

	return &mutex.Response{Granted: true}, nil
}

//func (mc *MutexClient) sendReleaseToPeer(peer string) {
//	conn, err := grpc.Dial(peer)
//	if err != nil {
//		log.Printf("Error connecting to peer %s: %v\n", peer, err)
//		return
//	}
//	defer conn.Close()
//
//	client := mutex.NewMutexClient(conn)
//
//	_, err = client.ReleaseAccess(context.Background(), &mutex.Request{ClientId: mc.clientId})
//	if err != nil {
//		log.Printf("Error releasing access to peer %s: %v\n", peer, err)
//	}
//}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run mutex_client.go <client-id> <peer1> <peer2> <peer3> ...")
	}

	clientId := os.Args[1]
	peers := os.Args[2:]

	time.Sleep(time.Second * 5)

	client := NewMutexClient(clientId, peers)
	client.askForAccess()

	log.Printf("%s had done it's task\n", clientId)

	time.Sleep(time.Second * 10)

}
