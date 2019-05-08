package main

import "fmt"

func main() {
	var x = []int{90, 15, 81, 87, 47, 59, 81, 18, 25, 40, 56, 8}

	i := 0
	l := len(x)
	for i < l {
		if x[i]%2 != 0 {
			x = append(x[:i], x[i+1:]...)
			l--
		} else {
			i++
		}
	}
	x = x[:i]

	fmt.Println(x)
}
