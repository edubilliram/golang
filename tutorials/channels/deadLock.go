//One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data.
//If this does not happen, then the program will panic at runtime with Deadlock.

//Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

package main

import "fmt"

/*func pross(ch chan string) {
	fmt.Println("write to channel")
	ch <- "GO"
}*/

func main() {
	ch := make(chan string)
	ch <- "GO"                         //// Panics here as it's not sure whether there is a reciver or not.
	fmt.Println("reciver here:", <-ch) /// even though we have a reciver here it panics in above line

}
