package main

import (
	"distributed_systems_mandatories/mandatory5/auction"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"slices"
	"time"
)

var highestBid int32 = 0
var highestBidderName string = "No bids yet"
var ongoing bool = true

var serverId = ""
var serverIp = ""
var timestamp int32 = 0
var waitingForAccess = false
var inCriticalSection = false
var replies = 0
var peerList = make([]string, 1)

type serverNode struct {
	auction.UnimplementedServerNodeServer
}

// RequestAccess this is the code that is responding to other serverNodes' requests
func (serverNode serverNode) Bid(ctx context.Context, request *auction.BidMessage) (*auction.Ack, error) {
	if !ongoing {
		return &auction.Ack{Outcome: "Fail"}, nil
	}

	//todo: make a thing that returns ack{exception}
	log.Printf("%v | %s made bid of %v", timestamp, request.Name, request.Amount)

	if request.Amount <= highestBid {
		return &auction.Ack{
			Outcome: "Fail",
		}, nil
	}

	bidSuccessful := serverNode.queueBid(*request)
	if bidSuccessful {
		return &auction.Ack{
			Outcome: "Success",
		}, nil
	}
	return &auction.Ack{
		Outcome: "Fail",
	}, nil

}

func (serverNode serverNode) Result(ctx context.Context, request *auction.Empty) (*auction.Outcome, error) {
	return &auction.Outcome{
		Ongoing: ongoing,
		Amount:  highestBid,
		Name:    highestBidderName,
	}, nil
}

// ask the other serverNodes for access
func (serverNode *serverNode) queueBid(bid auction.BidMessage) bool {
	log.Printf("%v | %s's peer list is: %v", timestamp, serverId, peerList)

	log.Printf("%v | %s is now trying to gain access to the critical section\n", timestamp, serverId)

	// get the highest timestamp and set this serverNode's timestamp to 1 higher
	for _, peer := range peerList {
		log.Printf("%v | Sending Lamport-timestamp request to %s\n", timestamp, peer)
		peerId, peerTimestamp := serverNode.sendTimestampRequestToPeer(peer)

		log.Printf("%v | %s's timestamp was: %v", timestamp, peerId, peerTimestamp)

		if peerTimestamp > timestamp {
			timestamp = peerTimestamp
		}
	}
	timestamp++
	waitingForAccess = true
	replies = 0

	for _, peer := range peerList {
		go serverNode.sendRequestToPeer(peer)
		log.Printf("%v | %s sent access request to %s\n", timestamp, serverId, peer)
	}

	// wait for all replies
	for replies < len(peerList) {
		time.Sleep(time.Millisecond * 100)
	}

	// Enter Critical Section
	log.Printf("%v | %s entered critical section\n", timestamp, serverId)
	inCriticalSection = true

	//access critical section
	if bid.Amount > highestBid {
		highestBid = bid.Amount
		highestBidderName = bid.Name

		// tell other server nodes that the auction has finished
		if timestamp >= 5 {
			ongoing = false
		}

		time.Sleep(time.Second * 5)
		serverNode.SendNewHighestBid(&bid, ongoing)

		inCriticalSection = false
		waitingForAccess = false
		return true
	}
	inCriticalSection = false
	waitingForAccess = false

	// Exit Critical Section
	log.Printf("%v | %s exited critical section\n", timestamp, serverId)
	return false
}

func (serverNode *serverNode) sendPeerListRequestToPeer(peer string) []string {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	// TODO: maybe don't create a new client? just use 'client'?
	requestingClient := auction.NewServerNodeClient(conn)

	response, err := requestingClient.RequestPeerList(context.Background(),
		&auction.Request{ServerNodeId: serverId, Timestamp: timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}

	return response.PeerList
}

func (serverNode *serverNode) sendTimestampRequestToPeer(peer string) (string, int32) {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)

	response, err := requestingClient.RequestLamportTimestamp(context.Background(),
		&auction.Request{ServerNodeId: serverId, Timestamp: timestamp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}
	return response.ServerNodeId, response.Timestamp
}

func (serverNode *serverNode) letPeerKnowIExist(peer string) {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
	}
	defer conn.Close()

	// TODO: maybe don't create a new client? just use 'client'?
	requestingClient := auction.NewServerNodeClient(conn)

	_, err = requestingClient.LetPeerKnowIExist(context.Background(),
		&auction.ServerNodeInfo{ServerNodeId: serverId, ServerNodeListeningIp: serverIp})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
	}
}

func (serverNode *serverNode) sendRequestToPeer(peer string) {
	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)

	_, err = requestingClient.RequestAccess(context.Background(), &auction.Request{
		ServerNodeId: serverId,
		Timestamp:    timestamp,
	})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
		return
	}

	log.Printf("%v | '%s' granted %s access to critical section\n", timestamp, peer, serverId)

	// increment replies if we got a reply from a client ahead of this client in the queue
	replies++
}

func (serverNode serverNode) SendNewHighestBid(bid *auction.BidMessage, ongoing bool) {
	for _, peer := range peerList {
		go func() {
			conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Error connecting to peer %s: %v", peer, err)
				return
			}
			defer conn.Close()

			requestingClient := auction.NewServerNodeClient(conn)

			_, err = requestingClient.ShareNewHighestBid(context.Background(), &auction.Outcome{
				Ongoing: ongoing,
				Amount:  bid.Amount,
				Name:    bid.Name,
			})
			if err != nil {
				log.Fatalf("Error requesting access from peer %s: %v", peer, err)
				return
			}

			log.Printf("%v | '%s' granted %s access to critical section\n", timestamp, peer, serverId)
		}()
	}
}

func (serverNode serverNode) ShareNewHighestBid(ctx context.Context, request *auction.Outcome) (*auction.Empty, error) {
	if !ongoing {
		return &auction.Empty{}, nil
	}

	highestBid = request.Amount
	highestBidderName = request.Name
	ongoing = request.Ongoing

	return &auction.Empty{}, nil
}

func (serverNode serverNode) LetPeerKnowIExist(ctx context.Context, request *auction.ServerNodeInfo) (*auction.Empty, error) {
	// add requesting client to peerList if it isn't already known
	if !slices.Contains(peerList, request.ServerNodeListeningIp) {
		peerList = append(peerList, request.ServerNodeListeningIp)

		log.Printf("%v | New peer added to %s's peer list: %s", timestamp, serverId, request.ServerNodeId)
	}

	log.Printf("%v | %s let %s know they exist", timestamp, request.ServerNodeId, serverId)
	return &auction.Empty{}, nil
}

// RequestLamportTimestamp returns this client's Lamport timestamp
func (serverNode serverNode) RequestLamportTimestamp(ctx context.Context, request *auction.Request) (*auction.LamportTimestamp, error) {
	log.Printf("%v | %s asked for %s's timestamp: %v", timestamp, request.ServerNodeId, serverId, timestamp)
	return &auction.LamportTimestamp{Timestamp: timestamp}, nil
}

// RequestPeerList returns this client's list of known peerList
func (serverNode serverNode) RequestPeerList(ctx context.Context, request *auction.Request) (*auction.PeerList, error) {
	log.Printf("%v | %s sent peer list to %s: %v", timestamp, serverId, request.ServerNodeId, peerList)
	return &auction.PeerList{PeerList: peerList}, nil
}

// RequestAccess this is the code that is responding to other serverNodes' requests
func (serverNode serverNode) RequestAccess(ctx context.Context, request *auction.Request) (*auction.Empty, error) {
	log.Printf("%v | %s asked %s for access to the critical section", timestamp, request.ServerNodeId, serverId)

	// wait to respond if the requesting serverNode is behind this serverNode in the queue
	if inCriticalSection || (waitingForAccess && request.Timestamp > timestamp) {
		for waitingForAccess || inCriticalSection {
			time.Sleep(time.Millisecond * 100)
		}
	}

	log.Printf("%v | %s granted %s access to the critical section", timestamp, serverId, request.ServerNodeId)

	return &auction.Empty{}, nil
}

func setUpListener(serverIp string) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()
	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", serverIp)
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	defer listener.Close()

	log.Printf("Started listening on ip: %s\n", serverIp)

	// Register the grpc server and serve its listener
	auction.RegisterServerNodeServer(grpcServer, &serverNode{})
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run server_node.go <server-id> <server-ip> <known-peer_server-ip>")
	}

	serverId = os.Args[1]
	serverIp = os.Args[2]

	if len(os.Args) > 3 {
		peerList = os.Args[3:]
	} else {
		peerList = make([]string, 0)
	}

	client := &serverNode{}
	log.Printf("%v | Created client: %s", timestamp, serverId)

	// if there is a known peer, ask that peer for the addresses of all the other peerList
	if len(os.Args) > 3 {
		knownPeersPeerList := client.sendPeerListRequestToPeer(peerList[0])
		peerList = append(peerList, knownPeersPeerList...)

		for _, peer := range peerList {
			client.letPeerKnowIExist(peer)
		}
		log.Printf("%v | %s let peers know they exist", timestamp, serverId)
	}

	// setup listener on client port
	go setUpListener(serverIp)

	time.Sleep(time.Hour)
}
