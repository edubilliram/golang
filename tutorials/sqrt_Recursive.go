package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) string {
	if n < 0 {
		res := sqrt(-n) + "i"
		return res
	}
	res := fmt.Sprint(math.Sqrt(n))
	return res
}

func main() {
	fmt.Println(sqrt(4))
}
