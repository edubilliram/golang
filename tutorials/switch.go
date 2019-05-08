package main

import (
	"fmt"
	"runtime"
)

func assert(cond bool, msg string) {
	if cond {
		fmt.Printf("asseration sussces: msg %s\n", msg)
	} else {
		fmt.Printf("asseration fail: msg %s\n", msg)
		panic(1)
	}
}
func main() {
	i5 := 5
	//i7 := 7
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.A")
	case "linux":
		fmt.Println("OS is Linux")
	default:
		fmt.Printf("Os is %s", os)
	}

	switch true {
	case i5 < 5:
		assert(false, "<")
	case i5 == 5:
		assert(true, "=")
	case i5 > 5:
		assert(false, ">")
	}

}
