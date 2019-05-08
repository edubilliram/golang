package main

import "fmt"

type details struct {
	ma   string
	node string
	hub  string
}

type allDetails struct {
	name         string
	relationship []details
}

func main() {
	a := allDetails{
		name:         "foo",
		relationship: []details{details{ma: "abc", node: "def", hub: "ghi"}, details{ma: "abc1", node: "def1", hub: "ghi1"}},
	}
	fmt.Println("a: ", a)
}
