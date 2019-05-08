package main

import "fmt"

func main() {
	v := 10
	incv(v)
	fmt.Println("value outside func", v)
	incp(&v)
	fmt.Println("pointer value outside func", v)

	va := [1]int{10}
	fArray(va)
	fmt.Println("array outside func", va)

	vs := []int{10}
	fSlice(vs)
	fmt.Println("slice outside func", vs)

	vm := map[int]int{0: 10}
	fMap(vm)
	fmt.Println("map outside func", vm)

	vstr := "10"
	fString(vstr)
	fmt.Println("string outside func", vstr)

}

func incv(a int) { // passed by value
	a = 99
	fmt.Println("value inside func", a)
}

func incp(a *int) { // passed by reference
	*a = 99
	fmt.Println("pointer value inside func", *a)
}

func fArray(a [1]int) { // arrays are passed by value by default
	a[0] = 99
	fmt.Println("array inside func", a)
}

func fSlice(a []int) { // slices are passed by reference by default
	a[0] = 99
	fmt.Println("slice inside func", a)
}

func fMap(a map[int]int) { // maps are passed by reference by default
	a[0] = 99
	fmt.Println("map inside func", a)
}

func fString(a string) {
	a = "99"
	fmt.Println("string inside func", a)
}

/*
value inside func 99
value outside func 10
pointer value inside func 99
pointer value outside func 99
array inside func [99]
array outside func [10]
slice inside func [99]
slice outside func [99]
map inside func map[0:99]
map outside func map[0:99]
string inside func 99
string outside func 10
*/
