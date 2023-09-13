// --- deadlock explanation ---
// The philosophers eat alternating between even i and odd i.
// Therefore, the odd philosophers will wait until the even has eaten.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the time each philosopher will delay picking up or putting down a fork.
const delayTime = 100

func main() {
	const philosopherCount = 5

	pickupLeftChannel := make(chan bool)
	putDownLeftChannel := make(chan bool)
	pickupLastChannel := pickupLeftChannel
	putDownLastChannel := putDownLeftChannel

	go forkThread(pickupLeftChannel, putDownLeftChannel)

	// an array we use to check if all philosophers has finished eating (3 times)
	var philosopherComplete = new([philosopherCount]chan bool)

	for i := 0; i < philosopherCount-1; i++ {
		philosopherComplete[i] = make(chan bool)

		pickupRightChannel := make(chan bool)
		putDownRightChannel := make(chan bool)

		go forkThread(pickupRightChannel, putDownRightChannel)
		go philosopherThread(i+1, philosopherComplete[i], pickupRightChannel, putDownRightChannel, pickupLeftChannel, putDownLeftChannel)

		pickupLeftChannel = pickupRightChannel
		putDownLeftChannel = putDownRightChannel
	}
	philosopherComplete[philosopherCount-1] = make(chan bool)

	go philosopherThread(philosopherCount, philosopherComplete[philosopherCount-1], pickupLastChannel, putDownLastChannel, pickupLeftChannel, putDownLeftChannel)

	// wait until all philosophers has eaten
	for i := 0; i < philosopherCount; i++ {
		<-philosopherComplete[i]
	}
	fmt.Println("All philosophers are finished eating!")

	time.Sleep(time.Millisecond * time.Duration(philosopherCount*delayTime*2))
}

func philosopherThread(i int, completeChannel chan bool, pickupRightChannel chan bool, putDownRightChannel chan bool, pickupLeftChannel chan bool, putDownLeftChannel chan bool) {
	eatCount := 0
	for eatCount < 3 {
		if i%2 == 0 {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
			pickupRightChannel <- true
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
			pickupLeftChannel <- true
		} else {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
			pickupLeftChannel <- true
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
			pickupRightChannel <- true
		}

		//time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
		//pickupRightChannel <- true
		//time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
		//pickupLeftChannel <- true

		eatCount++

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
		<-putDownRightChannel
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(delayTime)))
		<-putDownLeftChannel

		//fmt.Printf("Philosopher %d has eaten %d times\n", i, eatCount)
	}
	fmt.Printf("Philosopher %d is done eating!\n", i)

	completeChannel <- true
}

func forkThread(pickupChannel chan bool, putDownChannel chan bool) {
	for {
		<-pickupChannel
		putDownChannel <- true
	}
}
