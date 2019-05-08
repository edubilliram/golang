package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c1 := make(chan string)
	c2 := make(chan string)
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one" // write to a channel
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
		wg.Done()
	}()
	//fmt.Println("recived: %v\n", c1)
	// fmt.Println("recived: %v\n", c2)
	for i := 0; i < 2; i++ {
		// Await both of these values
		// simultaneously, printing each one as it arrives.
		select {
		case msg1 := <-c1: //// read from a channel
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	wg.Wait()

}
