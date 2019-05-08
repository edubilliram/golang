package main

import "fmt"

func digit(n int, ch chan int) {
	for n != 0 {
		ch <- n % 10
		n = n / 10
	}
	//	fmt.Println("channel:", <-ch)  //deadLock       // if we write from the channel here, it will be empty later at ln:19... In that case store in a varible and use it were ever it needed.
	close(ch) // closing channel should be used OR ELSE Deadlock , SO  while doing iterative things to unbufferedChannel ...... means while using "ok" OR "range" , for reading data from channels.
}

func calSqr(n int, sumCh chan int) {
	ch := make(chan int)
	sum := 0
	go digit(n, ch)
	for {
		v, ok := <-ch // if we don't close the channel still it tries to read , which gives deadlock
		if ok == true {
			sum += v * v
		} else {
			break
		}
	}
	sumCh <- sum

}
func calCub(n int, sumCh chan int) {
	ch := make(chan int)
	sum := 0
	go digit(n, ch)
	for v := range ch { // if we don't use "close", it will deadlock, as is still tries to read
		sum += v * v * v
	}
	sumCh <- sum
}
func main() {
	n := 123
	sqCh := make(chan int)
	cuCh := make(chan int)
	go calSqr(n, sqCh)
	go calCub(n, cuCh)

	fmt.Println("sum of squres:", <-sqCh+<-cuCh)
}
