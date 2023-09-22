package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const timeoutTime = 10

func main() {
	serverChannel := make(chan bool)
	clientChannel := make(chan bool)

	go server(serverChannel)
	go client(clientChannel)

	<-serverChannel
	<-clientChannel
}
func server(completeChannel chan bool) {
	wConn := createUDPWriter(8080, 8081)
	lConn := createUDPListener(8080)
	fmt.Println("Server listening on 127.0.0.1:8080")

	// read x
	x, remote := readIntFromConn(lConn)
	fmt.Printf("Server recieved: %d from %s\n", x, remote)

	// write x+1 and y
	data := []uint32{x + 1, 20}

	payload := make([]byte, 8)
	binary.BigEndian.PutUint32(payload[0:], data[0])
	binary.BigEndian.PutUint32(payload[4:], data[1])

	writeIntToConn(x+1, wConn)

	fmt.Println("Server finished!")
	completeChannel <- true
}

func client(completeChannel chan bool) {
	wConn := createUDPWriter(8081, 8080)
	lConn := createUDPListener(8081)

	// first send from client
	var x uint32 = 10
	byteArray := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteArray, x)
	writeIntToConn(byteArray, wConn)
	fmt.Printf("Client sent: %d\n", x)

	// client read x+1
	x1, _ := readIntFromConn(lConn)

	if x1 == x+1 {

	}

	fmt.Println("Client finished!")
	completeChannel <- true
}

func intsToByteArray(intArray []uint32) []byte {
	payload := make([]byte, len(intArray)*4) // Assuming int32 integers
	for i, num := range intArray {
		binary.BigEndian.PutUint32(payload[i*4:], num)
	}
	return payload
}

func byteArrayToInts(length int, byteArray []byte) []uint32 {
	var receivedData []uint32
	for i := 0; i < length; i += 4 {
		num := binary.BigEndian.Uint32(byteArray[i : i+4])
		receivedData = append(receivedData, num)
	}

	return receivedData
}

func writeIntToConn(byteArray []byte, wConn *net.UDPConn) {
	//bs := make([]byte, 4)
	//binary.LittleEndian.PutUint32(bs, i)
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
		panic(err)
	}

	// set a timeout deadline
	conn.SetDeadline(time.Now().Add(time.Second * timeoutTime))

	return conn
}
