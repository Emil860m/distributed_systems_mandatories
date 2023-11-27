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
var highestBidderName = "No bids yet"
var ongoing = true

var serverId = ""
var serverIp = ""
var timestamp int32 = 0
var bidTimestamp int32 = 0

var waitingForAccess = false
var inCriticalSection = false
var replies = 0
var peerIpList = make([]string, 1)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run server_node.go <server-id> <server-ip> <peer-server-ip-1> <peer-server-ip-2> ...")
	}

	serverId = os.Args[1]
	serverIp = os.Args[2]
	peerIpList = os.Args[3:]

	log.Printf("%v | Created server node: %s", timestamp, serverId)

	// setup listener on server IP and port
	go setUpListener(serverIp)

	time.Sleep(time.Hour)
}

type serverNode struct {
	auction.UnimplementedServerNodeServer
}

func (serverNode serverNode) Result(ctx context.Context, request *auction.Empty) (*auction.Outcome, error) {
	return &auction.Outcome{
		Ongoing: ongoing,
		Amount:  highestBid,
		Name:    highestBidderName,
	}, nil
}

func (serverNode serverNode) Bid(ctx context.Context, bid *auction.BidMessage) (*auction.Ack, error) {
	if !ongoing {
		return &auction.Ack{Outcome: "Fail: auction is finished"}, nil
	}

	log.Printf("%v | %s made bid of %v", timestamp, bid.Name, bid.Amount)

	if bid.Amount <= highestBid {
		return &auction.Ack{
			Outcome: "Fail: bid was not higher than highest bid",
		}, nil
	}

	// add bid to the shared bid-queue between server-nodes
	successfulBid := serverNode.queueBid(bid)

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

// queueBid Returns true if the bid was successfully placed
func (serverNode *serverNode) queueBid(bid *auction.BidMessage) bool {
	log.Printf("%v | %s is now trying to gain access to the critical section\n", timestamp, serverId)

	timestamp++
	bidTimestamp = timestamp

	bidSuccessful := false
	waitingForAccess = true
	replies = 0

	// ask other server-nodes for access critical section
	for _, peerIp := range peerIpList {
		go serverNode.sendAccessRequestToPeer(peerIp)
	}

	// wait for replies from all server-nodes in list of peers
	for replies < len(peerIpList) {
		time.Sleep(time.Millisecond * 100)
	}

	if !ongoing {
		waitingForAccess = false
		return false
	}

	// Enter critical section
	log.Printf("%v | %s entered critical section\n", timestamp, serverId)
	inCriticalSection = true

	time.Sleep(time.Second * 3) // access delay for easier manual testing

	if bid.Amount > highestBid {
		highestBid = bid.Amount
		highestBidderName = bid.Name

		log.Printf("%v | Sharing new highest bid: %v by %v", timestamp, bid.Amount, bid.Name)
		for _, peerIp := range peerIpList {
			serverNode.SendNewHighestBidToPeer(peerIp, bid, ongoing)
		}

		bidSuccessful = true
	}

	inCriticalSection = false
	waitingForAccess = false
	log.Printf("%v | %s exited critical section\n", timestamp, serverId)

	return bidSuccessful
}

func (serverNode *serverNode) sendAccessRequestToPeer(peerIp string) {
	log.Printf("%v | %s sent access request to %s\n", timestamp, serverId, peerIp)

	conn, err := grpc.Dial(peerIp, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error connecting to peerIp %s: %v", peerIp, err)
		removePeerIpFromList(peerIp)
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)

	_, err = requestingClient.RequestAccess(context.Background(), &auction.AccessRequest{
		ServerNodeId: serverId,
		BidTimestamp: bidTimestamp,
	})
	if err != nil {
		log.Printf("Error connecting to peerIp %s: %v", peerIp, err)
		removePeerIpFromList(peerIp)
		return
	}

	log.Printf("%v | '%s' granted %s access to critical section\n", timestamp, peerIp, serverId)

	// increment replies if we got a reply from another server ahead of this server in the bid queue
	replies++
}
func (serverNode serverNode) RequestAccess(ctx context.Context, accessRequest *auction.AccessRequest) (*auction.Empty, error) {
	log.Printf("%v | %s asked %s for access to the critical section", timestamp, accessRequest.ServerNodeId, serverId)

	// wait to respond if the requesting serverNode is behind this serverNode in the shared bid-queue
	if inCriticalSection || (waitingForAccess && accessRequest.BidTimestamp > bidTimestamp) {
		for (waitingForAccess || inCriticalSection) && ongoing {
			time.Sleep(time.Millisecond * 100)
		}
	}

	updateTimestamp(accessRequest.BidTimestamp)

	log.Printf("%v | %s granted %s access to the critical section", timestamp, serverId, accessRequest.ServerNodeId)

	return &auction.Empty{}, nil
}

func (serverNode serverNode) SendNewHighestBidToPeer(peerIp string, newHighestBid *auction.BidMessage, ongoing bool) {
	log.Printf("%v | Sharing new highest bid with %v", timestamp, peerIp)

	conn, err := grpc.Dial(peerIp, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to peerIp %s: %v", peerIp, err)
		return
	}
	defer conn.Close()

	requestingClient := auction.NewServerNodeClient(conn)

	_, err = requestingClient.ShareNewHighestBid(context.Background(), &auction.NewHighestBid{
		Amount:    newHighestBid.Amount,
		Name:      newHighestBid.Name,
		Timestamp: timestamp,
	})
	if err != nil {
		log.Fatalf("Error requesting access from peerIp %s: %v", peerIp, err)
		return
	}
}
func (serverNode serverNode) ShareNewHighestBid(ctx context.Context, newHighestBid *auction.NewHighestBid) (*auction.Empty, error) {
	if !ongoing {
		return &auction.Empty{}, nil
	}

	log.Printf("%v | Received new highest bid: %v by %v", timestamp, newHighestBid.Amount, newHighestBid.Name)

	highestBid = newHighestBid.Amount
	highestBidderName = newHighestBid.Name

	updateTimestamp(newHighestBid.Timestamp)
	if ongoing && timestamp >= auctionFinishTimestamp {
		ongoing = false
		log.Printf("%v | Auction finished! (%v >= %v)", timestamp, timestamp, auctionFinishTimestamp)
	}

	return &auction.Empty{}, nil
}

func setUpListener(serverIp string) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()
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
func removePeerIpFromList(peerIp string) {
	log.Printf("Removing peerIp '%s' from peerIp list", peerIp)
	for i := 0; i < len(peerIpList); i++ {
		if peerIpList[i] == peerIp {
			peerIpList[i] = peerIpList[len(peerIpList)-1]
			peerIpList = peerIpList[:len(peerIpList)-1]
			return
		}
	}
}
func updateTimestamp(newTimestamp int32) {
	if newTimestamp > timestamp {
		timestamp = newTimestamp
		log.Printf("%v | Timestamp updated", timestamp)
	}
}
