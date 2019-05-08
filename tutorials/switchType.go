package main

import "fmt"

func so(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("type interface:%T and value: %v\n", v, v)
	case string:
		fmt.Printf("type interface:%T and value: %v\n", v, v)
	default:
		fmt.Printf("type interface:%T and dont know about the value\n", v)
	}
}

func main() {
	so(21)
	so("me")
	so(false)
}
