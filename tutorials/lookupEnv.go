package main

import (
	"fmt"
	"os"
)

var (
	envSet  = "sample1"
	example = "env was set"
)

func read(envSet string, example string) string {
	//	_ = os.Setenv(envSet, "sample1")
	if _, eVar := os.LookupEnv(envSet); eVar {
		fmt.Printf("value set: %v\n", envSet)
		resp := fmt.Sprintln(example)
		return resp
	} else {
		resp := fmt.Sprintln("nothing set")
		return resp
	}
}

func main() {
	resp := read("sample1", example)
	fmt.Println(resp)
}
