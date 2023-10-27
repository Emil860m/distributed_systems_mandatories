package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// the max time the server and client will wait for a responce
const timeoutTime = 10

func main() {
	serverChannel := make(chan bool)
	clientChannel := make(chan bool)

	var serverPort = 8080
	var clientPort = 8081

	go server(serverChannel, serverPort)
	go client(clientChannel, serverPort, clientPort)

	// wait for both server and client to finish
	<-serverChannel
	<-clientChannel

	fmt.Println("\n3-way handshake finished!")
}
func server(completeChannel chan bool, serverPort int) {
	// server read x
	lConn := createUDPListener(serverPort)
	fmt.Printf("Server listening on 127.0.0.1:%d\n", serverPort)
	length, dataReceived, remote := readIntFromConn(lConn)
	lConn.Close()
	receivedIntArray := byteArrayToIntArray(length, dataReceived)
	x := receivedIntArray[0]

	fmt.Printf("Server received: syn seq=%d from %s\n", receivedIntArray[0], remote)

	// server send x+1, y
	var y = rand.Uint32()
	data := []uint32{receivedIntArray[0] + 1, y}
	wConn := createUDPWriter(serverPort, remote.Port, string(remote.IP))
	writeIntToConn(intArrayToByteArray(data), wConn)
	wConn.Close()

	fmt.Printf("Server sent: syn ack=%d, seq=%d\n", x+1, y)

	// server read y+1 x+1
	lConn = createUDPListener(serverPort)
	length, dataReceived, remote = readIntFromConn(lConn)
	lConn.Close()
	receivedIntArray = byteArrayToIntArray(length, dataReceived)

	fmt.Printf("Server received: ack=%d, seq=%d from %s\n", receivedIntArray[0], receivedIntArray[1], remote)
	if y+1 != receivedIntArray[0] || x+1 != receivedIntArray[1] {
		panic(fmt.Sprintf("Not correct response: y+1 != %d or x+1 != %d", receivedIntArray[0], receivedIntArray[1]))
	}

	fmt.Println("Server finished!")
	completeChannel <- true
}

func client(completeChannel chan bool, serverPort int, clientPort int) {
	time.Sleep(time.Millisecond * 200)

	// client send x
	var x = rand.Uint32()
	data := []uint32{x}
	byteArray := intArrayToByteArray(data)
	wConn := createUDPWriter(clientPort, serverPort, "127.0.0.1")
	writeIntToConn(byteArray, wConn)
	wConn.Close()

	fmt.Printf("Client sent: syn seq=%d\n", x)

	// client read x+1 y
	lConn := createUDPListener(clientPort)
	length, receivedByteArray, _ := readIntFromConn(lConn)
	lConn.Close()
	receivedIntArray := byteArrayToIntArray(length, receivedByteArray)

	fmt.Printf("Client received: syn ack=%d seq=%d\n", receivedIntArray[0], receivedIntArray[1])
	if x+1 != receivedIntArray[0] {
		panic(fmt.Sprintf("Not correct response: x+1 != %d", receivedIntArray[0]))
	}

	// client send y+1 x+1
	data = []uint32{receivedIntArray[1] + 1, x + 1}
	byteArray = intArrayToByteArray(data)
	wConn = createUDPWriter(clientPort, serverPort, "127.0.0.1")
	writeIntToConn(byteArray, wConn)
	wConn.Close()

	fmt.Printf("Client sent: ack=%d seq=%d\n", data[0], data[1])

	fmt.Println("Client finished!")
	completeChannel <- true
}

func intArrayToByteArray(intArray []uint32) []byte {
	payload := make([]byte, len(intArray)*4)
	for i, num := range intArray {
		binary.BigEndian.PutUint32(payload[i*4:], num)
	}
	return payload
}

func byteArrayToIntArray(length int, byteArray []byte) []uint32 {
	var receivedData []uint32
	for i := 0; i < length; i += 4 {
		num := binary.BigEndian.Uint32(byteArray[i : i+4])
		receivedData = append(receivedData, num)
	}

	return receivedData
}

func writeIntToConn(byteArray []byte, wConn *net.UDPConn) {
	_, writeErr := wConn.Write(byteArray)
	if writeErr != nil {
		fmt.Println("Write failed:", writeErr)
		panic(writeErr)
	}
}
func readIntFromConn(lConn *net.UDPConn) (int, []byte, *net.UDPAddr) {
	byteArray := make([]byte, 8)
	length, remote, readError := lConn.ReadFromUDP(byteArray[:])
	if readError != nil {
		fmt.Println("Read failed:", readError)
		panic(readError)
	}

	return length, byteArray, remote
}

func createUDPListener(port int) *net.UDPConn {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// set a timeout deadline
	conn.SetDeadline(time.Now().Add(time.Second * timeoutTime))

	return conn
}
func createUDPWriter(lPort int, rPort int, ip string) *net.UDPConn {
	raddr := net.UDPAddr{
		Port: rPort,
		IP:   net.ParseIP(ip),
	}
	laddr := net.UDPAddr{
		Port: lPort,
		IP:   net.ParseIP(ip),
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// set a timeout deadline
	conn.SetDeadline(time.Now().Add(time.Second * timeoutTime))

	return conn
}
