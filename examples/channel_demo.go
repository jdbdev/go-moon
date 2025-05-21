package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel that can hold strings
	messages := make(chan string)

	// Start a goroutine that sends a message after 2 seconds
	go func() {
		fmt.Println("Goroutine: I'm going to send a message in 2 seconds...")
		time.Sleep(2 * time.Second)
		messages <- "Hello from goroutine!" // Send message into channel
		fmt.Println("Goroutine: I sent the message!")
	}()

	fmt.Println("Main: Waiting for message...")
	msg := <-messages // Wait for and receive message from channel
	fmt.Println("Main: Got message:", msg)
}
