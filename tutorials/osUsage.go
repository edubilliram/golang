package main

import (
	"fmt"
	"io/ioutil"
)

func check(err error) error {
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

var a = []int{1, 2, 3, 4, 5, 6}

func main() {
	fmt.Printf("slice a: %v, len: %v, cap: %v\n", a, len(a), cap(a))
	b := a[1:4]
	fmt.Printf("slice b: %v, len: %v, cap: %v\n", b, len(b), cap(b))

	fmt.Println("-----------------------------------")
	dat, err := ioutil.ReadFile("/Users/V738294/go/src/tutorials/map.go")
	check(err)
	fmt.Println(string(dat))

}
