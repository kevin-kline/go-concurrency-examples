package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

// Uses Fan in function

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
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

func boring(msg string) <-chan Message { // Returns receive-only (<-) channel of Messages
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() { // Launch go routine inside function
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt // This line blocks everything until it receives true
		}
	}()
	return c // Returns channel to the caller
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true // This unblocks the boring for loop from Joe
		msg2.wait <- true // This unblocks the boring for loop from Ann
	}
	fmt.Println("You're both boring; I'm leaving.")
}
