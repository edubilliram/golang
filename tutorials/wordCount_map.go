package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	slcStr := strings.Split(s, " ")
	m := make(map[string]int)
	for _, v := range slcStr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
