package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("run run go routine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	//This line of code is blocking which means that until some Goroutine writes data to the done channel, the control will not move to the next line of code. Hence this eliminates the need for the time.Sleep
	fmt.Println(" excuting main function")
}
