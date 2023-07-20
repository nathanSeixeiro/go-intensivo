package main

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 11; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// thread 1
func main() {
	// go processando() // thread 2
	// processando()    // thread 1

	makeChannel()
}

func makeChannel() {
	channel := make(chan int)
	// anonymous function and reserved word 'go' means go routines
	go func() {
		// <- symbol n right osition means received the value (fill the channel)
		channel <- 1 // thread 2
	}()
	// <- symbol on left osition means unfill the channel
	fmt.Println(<-channel)
	time.Sleep(time.Second * 2)
}
