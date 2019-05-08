package main

import "fmt"

func main() {
	v := make(map[string]int)
	v["hey"] = 1
	v["gopher"] = 2
	fmt.Println("map : ", v)

	b := v["hey"]
	n := v["gopher"]
	fmt.Printf("let's print b and n : %v and  %v\n", b, n)
	fmt.Printf("let's print len : %v\n", len(v))

	delete(v, "gopher")
	fmt.Println(v)

	// using "_" gives bool dataType for chk OR it gives int dataType chk
	_, chk := v["gopher"]
	fmt.Println(chk)

	gopher := map[int]int{1: 2, 2: 3} // no need to give "make"
	fmt.Println(gopher)
}
