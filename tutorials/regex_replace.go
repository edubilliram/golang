package main

import (
	"fmt"
	"strings"
)

var replacer = strings.NewReplacer("!", " ", "@", " ", "#", " ", "$", " ", "%", " ", "&", " ", "*", " ", "  ", " ", ":", " ")

func modfyString(str string) string {
	str = replacer.Replace(str)
	return str
}

func main() {
	str := "aroot:*:0:0:System Administrator:/root:/bin/sh   replace   multiple   spaces   to   single   space"
	modfyStr := modfyString(str)
	fmt.Printf("modified string :%v\n", modfyStr)
}
