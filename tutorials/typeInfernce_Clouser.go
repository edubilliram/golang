package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int { // clouser for fibonacci series, aleays use func() as anony func
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()

	fmt.Println(f(), f(), f(), f(), f())
}
