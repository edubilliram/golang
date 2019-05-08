//Lets write one more program to understand channels better. This program will print the sum of the squares and cubes of the individual digits of a number.

//For example if 123 is the input, then this program will calculate the output as

//squares = (1 * 1) + (2 * 2) + (3 * 3)
//cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
//output = squares + cubes = 50

//Using Channels

package main

import (
	"fmt"
)

func calSquare(n int, sqChan chan int) {
	var m, s int
	for n != 0 {
		m = (n % 10) // modules operation --> gives reminder
		s += m * m
		n = n / 10 // divison operation -->quotient
	}
	sqChan <- s
}

func calCube(n int, cuChan chan int) {
	var m, s int
	for n != 0 {
		m = (n % 10)
		s += m * m * m
		n = n / 10
	}
	cuChan <- s
}

// without using channels
func compute(s []int) int {
	var sq, cu int
	for i := range s {
		sq += s[i] * s[i]

	}
	fmt.Printf("print square sum: %v\n", sq)
	for j := range s {
		cu += s[j] * s[j] * s[j]
	}
	fmt.Printf("print cube sum: %v\n", cu)
	return sq + cu
}

func main() {
	s := []int{1, 2, 3}
	a := compute(s)
	sqChan := make(chan int)
	cuChan := make(chan int)
	go calSquare(123, sqChan)
	go calCube(123, cuChan)
	b, c := <-sqChan, <-cuChan
	fmt.Println("compute: ", a)
	fmt.Println("calSquare:", b)
	fmt.Println("calCube:", c)
	fmt.Println("sum:", b+c)
}
