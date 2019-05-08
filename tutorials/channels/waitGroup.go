package main

import (
	"fmt"
	"sync"
)

func add(slc [][]int) (result []int) {
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, v := range slc {

			wg.Add(1)
			go func(v1 []int) {

				defer wg.Done()
				sum := 0
				for _, l := range v1 {
					sum += l
				}
				fmt.Println("sum: ", sum)
				ch <- sum
			}(v)

		}

		wg.Wait()

		close(ch)
	}()
	for s := range ch {
		result = append(result, s)
	}

	return result
}

func main() {

	a := [][]int{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}
	res := add(a)

	fmt.Printf("read result: %v,  type: %T\n", res, res)

}
