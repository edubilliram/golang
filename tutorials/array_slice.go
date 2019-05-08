package main

import (
	"fmt"
)

var dy = [6]string{"a", "b", "c", "d", "e", "f"} //array declare and assign
var dx []string

func main() {
	fmt.Printf("%q\n", dy)
	a := [...]string{"Penn", "Teller", "gopher"} //array declare and assign
	fmt.Println("array a :", a)

	letters := []string{"a", "b", "c", "d", "e", "f"} //slice
	fmt.Printf("letters slice: %q\n", letters)

	s := make([]string, 3, 10) // make buit-in func "func make([]T, len, cap) []T"
	fmt.Printf("s slice: %q\n", s)

	var sb []byte           // create a variable of type byte
	sb = make([]byte, 5, 5) // make a slice of that variable with len and cap
	fmt.Printf("sb slice: %v\n", sb)

	sbr := []byte{1, 1, 1, 1} //intialize and assign slice
	fmt.Printf("sbr :%v\n", sbr)

	//nlslc := []int
	var nlslc []int // nil clice of len=0 and cap =0
	fmt.Printf("nlslc :%v\n", nlslc)

	nlslc1 := make([]string, 5) // make empty slices of len 5
	fmt.Printf("nlslc1 :%v\n", nlslc1)

	//Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'} //slicing
	fmt.Printf("slicing b: %s\n", b[1:4])
	fmt.Printf("slicing b: %s\n", b[:2])
	fmt.Printf("slicing b: %s\n", b[2:])
	fmt.Printf("slicing b: %s\n", b[:])
	c := b[:] // or c := b
	fmt.Printf("slice c: %s\n", c[:])
	d := b[:cap(b)]
	fmt.Printf("slice d: %s\n", d)

	// create a slice of double cap of b
	t := make([]byte, len(b), (cap(b)+1)*2) // +1 in case cap(s) == 0
	for i := range b {
		t[i] = b[i]
	}
	fmt.Printf("slice t: %q\n", t)

}
