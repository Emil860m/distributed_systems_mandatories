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
	lConn := createUDPListener(8080)
	defer lConn.Close()
	fmt.Println("Server listening on 127.0.0.1:8080")
	// read dataReceived
	length, dataReceived, remote := readIntFromConn(lConn)
	receivedIntArray := byteArrayToInts(length, dataReceived)
	x := receivedIntArray[0]
	fmt.Printf("Server recieved: %d from %s\n", receivedIntArray, remote)
	lConn.Close()

	// send dataReceived+1, y
	var y uint32 = 20
	data := []uint32{receivedIntArray[0] + 1, y}
	wConn := createUDPWriter(8080, remote.Port, string(remote.IP))
	defer wConn.Close()
	writeIntToConn(intsToByteArray(data), wConn)
	fmt.Printf("Server sent: %d, %d\n", x+1, y)
	wConn.Close()

	lConn = createUDPListener(8080)
	// read dataReceived
	length, dataReceived, remote = readIntFromConn(lConn)
	receivedIntArray = byteArrayToInts(length, dataReceived)
	fmt.Printf("Server recieved: %d from %s\n", receivedIntArray, remote)
	if y+1 != receivedIntArray[0] || x+1 != receivedIntArray[1] {
		panic("Not correct response")
	}
	lConn.Close()

	fmt.Println("Server finished!")
	completeChannel <- true
}

func client(completeChannel chan bool) {
	time.Sleep(5000000000000)
	// first send from client
	wConn := createUDPWriter(8081, 8080, "127.0.0.1")
	defer wConn.Close()
	var x uint32 = 10
	data := []uint32{10}
	byteArray := intsToByteArray(data)
	writeIntToConn(byteArray, wConn)
	fmt.Printf("Client sent: %d\n", x)
	wConn.Close()

	// client read x+1
	lConn := createUDPListener(8081)
	defer lConn.Close()
	length, receivedByteArray, _ := readIntFromConn(lConn)
	lConn.Close()
	receivedIntArray := byteArrayToInts(length, receivedByteArray)
	fmt.Printf("Client recieved: %d\n", receivedIntArray)
	if x+1 != receivedIntArray[0] {
		panic("Not correct response")
	}

	wConn = createUDPWriter(8081, 8080, "127.0.0.1")
	data = []uint32{receivedIntArray[1] + 1, x + 1}
	byteArray = intsToByteArray(data)
	writeIntToConn(byteArray, wConn)
	fmt.Printf("Client sent: %d, %d\n", data[0], data[1])
	wConn.Close()

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
