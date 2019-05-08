package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func set() int {
	vals := &Vertex{X: 2, Y: 3}
	return vals.X + vals.Y
}

func set1(i interface{}) int {

}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X)
	v.X = 4
	fmt.Println(v.X)
	a := set()
	fmt.Println("a:", a)
}
