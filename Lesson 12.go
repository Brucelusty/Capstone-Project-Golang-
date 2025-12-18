// ============================================================================
// LESSON 12: CHANNELS (COMMUNICATION BETWEEN GOROUTINES)
// ============================================================================

package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to channel
		fmt.Printf("Sent: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close channel when done
}

func lesson12_Channels() {
	fmt.Println("\n=== LESSON 12: CHANNELS ===")
	
	// Create a channel
	ch := make(chan int)
	
	// Start goroutine to send data
	go sendData(ch)
	
	// Receive data from channel
	for value := range ch { // Loop until channel is closed
		fmt.Printf("Received: %d\n", value)
	}
	
	fmt.Println("Channel closed, all data received!")
	
	// Buffered channel (can hold multiple values)
	bufferedCh := make(chan string, 3)
	bufferedCh <- "First"
	bufferedCh <- "Second"
	bufferedCh <- "Third"
	
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
}