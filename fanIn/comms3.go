package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Uses Fan in function

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

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
	c := fanIn(boring("Joe"), boring("Ann"))

	// joe := boring("Joe")
	// ann := boring("Ann")

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
