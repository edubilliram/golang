package main

import "fmt"

var m = make(map[int]int) // nil map, writing to nil map always gives panic during runtime

func main() {
	//m = make(map[int]int) //

	m[1] = 2
	fmt.Println(m)
}
