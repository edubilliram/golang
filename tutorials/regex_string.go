package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "aroot:*:0:0:System Administrator:/root:/bin/sh   replace   multiple   spaces   to   single   space"
	modStr := nAllNumReplacer(str)
	fmt.Printf("modifed String str: %v\n", modStr)
}

func nAllNumReplacer(str string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Printf("invalid reg :%v\n", err)
	}
	modStr := reg.ReplaceAllString(str, " ")
	return modStr
}
