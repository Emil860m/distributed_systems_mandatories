package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var philosopherCount = 5

func philosopherLogic(i int, pickUpRight chan int, putDownRight chan int, pickUpLeft chan int, putDownLeft chan int) {
	nomnom := 0
	for nomnom < 3 {
		if i%2 == 0 {
			pickUpRight <- i
			pickUpLeft <- i
		} else {
			pickUpLeft <- i
			pickUpRight <- i
		}
		<-putDownLeft
		<-putDownRight
		nomnom++
		fmt.Println("philosopher " + strconv.Itoa(i) + " has eaten " + strconv.Itoa(nomnom) + " times")
	}
}

func forkLogic(i int, pickUp chan int, putDown chan int) {
	for {
		<-pickUp
		putDown <- i
	}
}

func philosopherRandomizerLogic(pickUpRight chan int, putDownRight chan int, pickUpLeft chan int, putDownLeft chan int) {
	nomnom := 0
	for nomnom < 3 {
		if rand.Intn(100000000) == 1 {
			pickUpRight <- 1
			pickUpLeft <- 1
			<-putDownLeft
			<-putDownRight
			nomnom++
			fmt.Println("philosopher has eaten " + strconv.Itoa(nomnom) + " times")
		}
	}
}

func main() {
	lastChan1 := make(chan int)
	lastChan2 := make(chan int)
	previousChan1 := lastChan1
	previousChan2 := lastChan2
	go forkLogic(philosopherCount-1, lastChan1, lastChan2)
	for i := 0; i < philosopherCount-1; i++ {
		c1 := make(chan int)
		c2 := make(chan int)
		go forkLogic(i, c1, c2)
		//go philosopherLogic(i, c1, c2, previousChan1, previousChan2)
		go philosopherRandomizerLogic(c1, c2, previousChan1, previousChan2)
		previousChan1 = c1
		previousChan2 = c2
	}
	//go philosopherLogic(philosopherCount-1, lastChan1, lastChan2, previousChan1, previousChan2)
	go philosopherRandomizerLogic(lastChan1, lastChan2, previousChan1, previousChan2)
	time.Sleep(10_000_000_000)
}
