package main

import "fmt"

func main() {
	sum := 1
	// here for loop iterates untill sum < 1000 and then terminates the loop ==sum = 1024
	// if it's if loop it only checks once == sum =2
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
