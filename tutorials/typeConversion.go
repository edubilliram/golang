package main

import "fmt"
import "strconv"

func main() {
	var a, b int = 1, 2
	//	var c string = string(a) for anytype to string we have import a "strconv" packg
	var c string = strconv.Itoa(a)
	fmt.Printf("%T\t,c= %v\n", c, c)
	fmt.Println(a, b, c)
}
