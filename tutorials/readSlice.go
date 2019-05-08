package main

import "fmt"

func main() {

	var csvData = [][]string{}

	csvData = [][]string{
		[]string{"k1", "v11", "v12"},
		[]string{"k1", "v21", "v22"},
		[]string{"k2", "v21", "v22"},
		[]string{"k3", "v31", "v32"},
	}

	for _, c := range csvData {
		fmt.Printf("internal slice: %v\n", c)
		for _, d := range c {
			fmt.Printf("internal slice: %v\n", d)
		}
	}

	fmt.Println(csvData)
}
