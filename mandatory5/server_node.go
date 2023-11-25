package main

import (
	"distributed_systems_mandatories/mandatory5/auction"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"time"
)

const auctionFinishTimestamp = 5

var highestBid int32 = 0
var highestBidderName string = "No bids yet"
var ongoing bool = true

var serverId = ""
var serverIp = ""
var timestamp int32 = 0
var bidTimestamp int32 = 0
var waitingForAccess = false
var inCriticalSection = false
var replies = 0
var peerList = make([]string, 1)

type serverNode struct {
	auction.UnimplementedServerNodeServer
}

//TODO: Fault tolerance
//todo: remove server from peerList if they don't respond (maybe set timeout to 1 second and then remove the peer when err != nil)
//todo: maybe add server to peerList when an unknown server asks us for access
//todo: read the fault tolerance description on the LearnIT page

func (serverNode serverNode) Bid(ctx context.Context, request *auction.BidMessage) (*auction.Ack, error) {
	if !ongoing {
		return &auction.Ack{Outcome: "Fail: auction is finished"}, nil
	}

	//todo: make a thing that returns ack{exception}
	log.Printf("%v | %s made bid of %v", timestamp, request.Name, request.Amount)

	if request.Amount <= highestBid {
		return &auction.Ack{
			Outcome: "Fail: bid was not higher than highest bid",
		}, nil
	}

	successfulBid := serverNode.queueBid(request)

	if ongoing && timestamp >= auctionFinishTimestamp {
		ongoing = false
		log.Printf("%v | Auction finished! (%v >= %v)", timestamp, timestamp, auctionFinishTimestamp)
	}

	if successfulBid {
		return &auction.Ack{
			Outcome: "Success",
		}, nil
	}

	if !ongoing {
		return &auction.Ack{Outcome: "Fail: auction is finished"}, nil
	}

	return &auction.Ack{
		Outcome: "Fail: bid was not higher than highest bid",
	}, nil

}

func (serverNode serverNode) Result(ctx context.Context, request *auction.Empty) (*auction.Outcome, error) {
	return &auction.Outcome{
		Ongoing: ongoing,
		Amount:  highestBid,
		Name:    highestBidderName,
	}, nil
}

func (serverNode *serverNode) queueBid(bid *auction.BidMessage) bool {
	bidSuccessful := false

	log.Printf("%v | %s's peer list is: %v", timestamp, serverId, peerList)

	log.Printf("%v | %s is now trying to gain access to the critical section\n", timestamp, serverId)

	timestamp++
	bidTimestamp = timestamp
	waitingForAccess = true
	replies = 0

	for _, peer := range peerList {
		go serverNode.sendAccessRequestToPeer(peer)
	}

	// wait for all replies
	for replies < len(peerList) {
		time.Sleep(time.Millisecond * 100)
	}
	if !ongoing {
		waitingForAccess = false
		return false
	}

	// Enter Critical Section
	log.Printf("%v | %s entered critical section\n", timestamp, serverId)
	inCriticalSection = true

	//access critical section
	time.Sleep(time.Second * 3) // access delay for easier manual testing

	if bid.Amount > highestBid {
		highestBid = bid.Amount
		highestBidderName = bid.Name

		log.Printf("%v | Sharing new highest bid: %v by %v", timestamp, bid.Amount, bid.Name)
		for _, peer := range peerList {
			serverNode.SendNewHighestBid(peer, bid, ongoing)
		}

		bidSuccessful = true
	}

	inCriticalSection = false
	waitingForAccess = false
	log.Printf("%v | %s exited critical section\n", timestamp, serverId)

	return bidSuccessful
}

func (serverNode *serverNode) sendAccessRequestToPeer(peer string) {
	log.Printf("%v | %s sent access request to %s\n", timestamp, serverId, peer)

	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)

	_, err = requestingClient.RequestAccess(context.Background(), &auction.AccessRequest{
		ServerNodeId: serverId,
		BidTimestamp: bidTimestamp,
	})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
		return
	}

	log.Printf("%v | '%s' granted %s access to critical section\n", timestamp, peer, serverId)

	// increment replies if we got a reply from another server ahead of this server in the bid queue
	replies++
}

// RequestAccess this is the code that is responding to other serverNodes' requests
func (serverNode serverNode) RequestAccess(ctx context.Context, request *auction.AccessRequest) (*auction.Empty, error) {
	log.Printf("%v | %s asked %s for access to the critical section", timestamp, request.ServerNodeId, serverId)

	// wait to respond if the requesting serverNode is behind this serverNode in the queue
	if inCriticalSection || (waitingForAccess && request.BidTimestamp > bidTimestamp) {
		for (waitingForAccess || inCriticalSection) && ongoing {
			time.Sleep(time.Millisecond * 100)
		}
	}

	updateTimestamp(request.BidTimestamp)

	log.Printf("%v | %s granted %s access to the critical section", timestamp, serverId, request.ServerNodeId)

	return &auction.Empty{}, nil
}

func (serverNode serverNode) SendNewHighestBid(peer string, bid *auction.BidMessage, ongoing bool) {
	log.Printf("%v | Sharing new highest bid with %v", timestamp, peer)

	conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peer %s: %v", peer, err)
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)

	_, err = requestingClient.ShareNewHighestBid(context.Background(), &auction.NewHighestBid{
		Amount:    bid.Amount,
		Name:      bid.Name,
		Timestamp: timestamp,
	})
	if err != nil {
		log.Fatalf("Error requesting access from peer %s: %v", peer, err)
		return
	}

	log.Printf("%v | '%s' granted %s access to critical section\n", timestamp, peer, serverId)
}

func (serverNode serverNode) ShareNewHighestBid(ctx context.Context, request *auction.NewHighestBid) (*auction.Empty, error) {
	if !ongoing {
		return &auction.Empty{}, nil
	}

	log.Printf("%v | Received new highest bid: %v by %v", timestamp, request.Amount, request.Name)

	highestBid = request.Amount
	highestBidderName = request.Name

	updateTimestamp(request.Timestamp)
	if ongoing && timestamp >= auctionFinishTimestamp {
		ongoing = false
		log.Printf("%v | Auction finished! (%v >= %v)", timestamp, timestamp, auctionFinishTimestamp)
	}

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

	log.Printf("%v | Started listening on ip: %s\n", timestamp, serverIp)

	// Register the grpc server and serve its listener
	auction.RegisterServerNodeServer(grpcServer, &serverNode{})
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func updateTimestamp(newTimestamp int32) {
	if newTimestamp > timestamp {
		timestamp = newTimestamp
		log.Printf("%v | Timestamp updated", timestamp)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run server_node.go <server-id> <server-ip> <peer-server-ip-1> <peer-server-ip-2> ...")
	}

	serverId = os.Args[1]
	serverIp = os.Args[2]
	peerList = os.Args[3:]

	log.Printf("%v | Created server node: %s", timestamp, serverId)

	// setup listener on server IP and port
	go setUpListener(serverIp)

	time.Sleep(time.Hour)
}
