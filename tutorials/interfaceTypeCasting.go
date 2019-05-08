package main

import "fmt"

func main() {
	var d interface{} = "hello"

	s := d.(string)
	fmt.Println(s)

	p, ok := d.(float64) // it panics if it's wrongly typecasteing
	if !ok {
		fmt.Println(p, ok)
	}

}
