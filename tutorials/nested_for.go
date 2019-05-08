package main

import "fmt"

func compare(aa, bb []int) []int {
	var out []int
	//bioCount = num of CPU cycles
	bigoCount := 0
	for _, a := range aa {
		for _, b := range bb {
			bigoCount++
			if a == b {
				out = append(out, a)
				continue
			}
		}
	}
	fmt.Println("bigoCount:", bigoCount)
	return out
}

func main() {

	aa := []int{1, 2, 3, 4, 5, 6, 7}
	bb := []int{4, 7, 9, 0, 1, 32, 5, 6}
	out := compare(aa, bb)
	fmt.Printf("common num's stored in out : %v\n", out)
}
