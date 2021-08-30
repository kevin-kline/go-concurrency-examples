package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creates channel inside function
// Runs goroutine inside function

func boring(msg string, quit chan bool) <-chan string { // Returns receive-only (<-) channel of strings
	c := make(chan string)
	go func() { // Launch go routine inside function
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Returns channel to the caller
}

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	// quit after some executions
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	// There is a problem if they need to do something, so quit channel...
}
