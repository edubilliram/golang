package main

import (
	"fmt"
	"math"
)

func pwr(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {

	fmt.Println(pwr(2, 3, 10))
}
