package main

import (
	"flag"
	"fmt"
)

func main() {
	var test string
	flag.StringVar(&test, "test", "dev", "usage:prod/qa/dev")
	flag.Parse()
	fmt.Println(test)
}
