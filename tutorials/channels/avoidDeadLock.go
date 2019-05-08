//One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data.
//If this does not happen, then the program will panic at runtime with Deadlock.

//Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() { /// It's a separate process begins because of go routine, So here Reciver in ln:21 will be ready, by the time something written to channel
		fmt.Println("separate go routine to write to channel: ch")
		ch <- "GO"
	}()
	fmt.Println("reciver here for ch:", <-ch)

}
