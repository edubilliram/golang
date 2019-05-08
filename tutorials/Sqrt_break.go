package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {

		z -= (z*z - x) / (2 * z)
		fmt.Printf("value of z, z*z :%v, %v\n ", z, z*z)
		if math.Sqrt(x) == z {
			fmt.Printf("itration :%v\n", i)
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), Sqrt(9))
}

//https://tour.golang.org/flowcontrol/8
