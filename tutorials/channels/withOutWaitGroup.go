package main

import (
	"fmt"
	"time"
)

func add(slc [][]int) (result []int) {
	ch := make(chan int)
	for _, v := range slc {

		go func(v1 []int) {
			sum := 0
			for _, l := range v1 {
				sum += l
			}
			ch <- sum
		}(v)

	}

	go func() {
		for s := range ch {
			result = append(result, s)
		}
	}()
	time.Sleep(2 * time.Second)
	return result
}

func main() {

	a := [][]int{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}
	res := add(a)

	fmt.Printf("read result: %v,  type: %T\n", res, res)

}
