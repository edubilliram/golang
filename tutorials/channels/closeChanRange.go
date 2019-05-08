package main

import (
	"fmt"
)

func proc(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go proc(ch)
	for v := range ch {
		fmt.Println("value recived:", v)
	}
}
