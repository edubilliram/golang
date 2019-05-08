package main

import "fmt"

func main() {
	ch := make(chan int, 2) // try making unbuffered --> gives deadlock---> to over come this ln:7 should be goRoutine OR ch should be Buffered
	ch <- 1                 // it will be deadlock here, if ch is not buffered channel... because unbuff channel don't store values in any buffer/ address loc and don't wait for reciver
	ch <- 2
	fmt.Println("reciver ready:", <-ch)
	fmt.Println("reciver ready:", <-ch)
}
