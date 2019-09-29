package Channels

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// Channels can hold two values
// The value of the channel and its boolean
func SimpleChannels() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channels"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}

// Closeing channels and
// Useing the range keyword -- supports channels as arguments and will break the loop when a channel
// Is closed
func FuncOne() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Printf("\n%v ", integer)
	}
}

// Buffer Channels
func FuncTwo() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo((os.Stdout))

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v\n", integer)
	}
}

// Simple select channels
func FuncThree() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep((5 * time.Second))
		close(c)
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later \n", time.Since(start))
	}
}
