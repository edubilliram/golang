package main

import (
	"fmt"
	"regexp"
	"strings"
)

var s []string

func cleanString(str string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Printf("invalid reg :%v\n", err)
	}
	modStr := reg.ReplaceAllString(str, " ")
	return modStr
}

func split(str string) []string {
	str = cleanString(str)
	//	modSlc := strings.Split(str, " ")
	//	fmt.Printf("modified slice :%v\n", modSlc)
	// return modSlc
	return strings.Split(str, " ")
}

func group(s) (modSlc [][]string) {

}

func main() {
	s = "Today we *celebrate*   $8 years since @Go was released as an open-source project@"
	prosSlc := split(s)
	fmt.Printf("processed string as slice :%q\n", prosSlc)

}
