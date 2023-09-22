package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	serverChannel := make(chan bool)
	clientChannel := make(chan bool)

	go server(serverChannel)
	go client(clientChannel)

	<-serverChannel
	<-clientChannel
}
func server(completeChannel chan bool) {
	lConn := createUDPListener(8080)
	fmt.Println("Server listening on 127.0.0.1:8080")

	x, remote := receiveIntFromConn(lConn)
	fmt.Printf("Server recieved: %d from %s\n", x, remote)

	fmt.Println("Server finished!")
	completeChannel <- true
}

func client(completeChannel chan bool) {
	wConn := createUDPWriter(8081, 8080)
	lConn := createUDPListener(8081)

	// first send from client
	var x uint32 = 10

	sendIntToConn(x, wConn)
	fmt.Printf("Client sent: %d\n", x)

	x1, remote := receiveIntFromConn(lConn)

	if x1 == x+1 {

	}

	fmt.Println("Client finished!")
	completeChannel <- true
}

func sendIntToConn(i uint32, wConn *net.UDPConn) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, i)
	_, writeErr := wConn.Write(bs)
	if writeErr != nil {
		fmt.Println("failed:", writeErr)
		return
	}
}
func receiveIntFromConn(lConn *net.UDPConn) (uint32, *net.UDPAddr) {
	bs := make([]byte, 4)
	_, remote, _ := lConn.ReadFromUDP(bs[:])
	x := binary.LittleEndian.Uint32(bs)

	return x, remote
}

func createUDPListener(port int) *net.UDPConn {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return conn
}
func createUDPWriter(lPort int, rPort int) *net.UDPConn {
	raddr := net.UDPAddr{
		Port: rPort,
		IP:   net.ParseIP("127.0.0.1"),
	}
	laddr := net.UDPAddr{
		Port: lPort,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return conn
}
