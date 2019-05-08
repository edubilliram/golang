package main

import "fmt"

var slc, dop []int

func main() {
	slc := []int{1, 2, 3, 4, 5, 5}
	for _, v := range slc {
		//fmt.Println("values: ", v)
		for _, d := range dop {
			if d == v {
				fmt.Println("common: ", d)
			}
		}
		dop = append(dop, v)
	}
	fmt.Println(dop)
}
