package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Print() {
	fmt.Printf("vertex: {x: %v, y: %v}\n\n", v.X, v.Y)
}

func (v *Vertex) square() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
	v.Print()
}

func main() {
	v := &Vertex{3, 4}
	v.Print()

	v.square()

	v.Print()

}
