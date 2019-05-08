package main

import "fmt"

func proc(ch chan<- int) {
	fmt.Println("ch is an unidirectional channel")
	ch <- 10
}

/*
func main() {
	ch := make(chan<- int)
	go proc(ch)
	///////////////ERROR : ./uniDirChan.go:14:15: invalid operation: <-ch (receive from send-only type chan<- int)  // becasue it's unidirectional
	fmt.Println(<-ch)
}*/
func main() {
	ch := make(chan int) //this is bidirectional
	go proc(ch)          //uniDir chan is converted to BiDir chan here
	fmt.Println(<-ch)
}
