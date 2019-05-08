package main

import (
	"fmt"
	"os"
)

func main() {
	/*for i, a := range os.Args[1:] {
	fmt.Printf("Argument %d is %s\n", i+1, a)
	}
	*/
	// always the first Arugmunt is Program Command (implicitly), here "./osArgs"
	//////////////

	/*var s = os.Args[:]
	fmt.Printf("Argument %s\n", s)*/
	//////////////

	var s, sep string 
	fmt.Printf("type: %T and value: %v\n",os.Args, len(os.Args))

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// go run osArgs.go 1 a 3 4 5 [1,2,3]

//to make it more simpler you can use  strings.Join()  to print os.Ags[1:]
