package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes
	big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	small = big >> 99
)

func printInt(x int) int           { return x*10 + 1 }
func printFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(printInt(small))
	fmt.Println(printFloat(big))
	fmt.Println(printFloat(small))
}
