package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int { // clouser func
		temp := x
		x = y
		y = temp + y
		return temp // check diff return variables
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
