package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creates channel inside function
// Runs goroutine inside function

func boring(msg string) <-chan string { // Returns receive-only (<-) channel of strings
	c := make(chan string)
	go func() { // Launch go routine inside function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Returns channel to the caller
}

func main() {
	//c := boring("boring!") // Returns a channel
	joe := boring("Joe")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		// fmt.Printf("You say: %q\n", <-c)
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring; I'm leaving.")
}
