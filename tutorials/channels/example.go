package main

import "fmt"

func sendChan(ch chan<- int) {
	ch <- 10
}

/// verify unidirectional channel
/*func main() {
	ch := make(chan<- int)
	go sendChan(ch)
	fmt.Println(<-ch) /// verify unidirectional channel
}*/

// convert uidirectional chan to biderctional chan
/*func main() {
	ch := make(chan int)
	go sendChan(ch)
	fmt.Println(<-ch)
}*/
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

//Check closing channels
/*func main() {
	ch := make(chan int)
	go producer(ch)
	for { // infinate for loop helps in checking v continously, with out for loop ... we will check only once
		v, ok := <-ch
		if ok == false {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("value recived from ch", v, ok)
	}
}*/

//closing channels using: for-range, recives values until channel closed
/*func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("recived v:", v)
	}
}*/

//writing a func to use as a concurrent go routine
func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("writing succesfully to ch")
	}
	close(ch)
}

// check, blocking using BUFFERED channels
func main() {
	ch := make(chan int, 5)
	go write(ch)
	for v := range ch {
		fmt.Println("recived:", v)
	}
}
