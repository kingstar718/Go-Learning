package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Send a value to notify that we're done
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)
	// Block util we receive a notification from the worker on the channel
	<-done
}
