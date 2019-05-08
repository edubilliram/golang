package main

import "fmt"

var pow = make([]int, 10)

func main() {
	pow1 := []int{1, 2, 4, 6, 8, 10, 12, 14}
	for i, _ := range pow {
		pow[i] = 1 << uint(i) // full equation: (x << n == x*2^n ) (x >> n == x*2^(-n))
		fmt.Printf("%v\n", pow[i])
	}
	for i, v := range pow1 {
		pow1[i] = v << uint(i) // full equation: (x << n == x*2^n ) (x >> n == x*2^(-n))
		//fmt.Printf("%v\n", pow1[i])
	}
	fmt.Printf("%v\n", pow1)
	fmt.Printf("%v\n", pow)
}
