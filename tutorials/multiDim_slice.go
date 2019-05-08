package main

import (
	"fmt"
)

func main() {
	a := [][]int{[]int{1, 2, 3}, []int{3, 4, 5}, []int{5, 6, 7}}

	fmt.Println(a[0])
	fmt.Println(a[0][0])

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("%v", a[i][j])
		}
		fmt.Print("\n")
	}
}
