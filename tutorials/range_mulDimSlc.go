package main

import "fmt"

var pow = []int{1, 2, 4, 6, 8, 10, 12, 14}
var slc = [][]int{[]int{1, 2, 3}, []int{3, 4, 5}, []int{5, 6, 7}}

func main() {
	for i, v := range pow {
		fmt.Printf("for %v - 2*%v = %v\n", i, i, v)
	}

	for a := range pow {
		fmt.Println("default always returns INDEX value of range not VALUE in  slice", a)
	}

	for i, _ := range slc {
		for j, _ := range slc[i] {
			fmt.Printf("%v ", slc[i][j])
		}
		fmt.Print("\n")
	}

}
