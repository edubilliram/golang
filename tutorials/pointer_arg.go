package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func Print(v Vertex) {
	fmt.Printf("vertex: {x: %v, y: %v}\n\n", v.X, v.Y)
}

func square(v *Vertex) {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
	Print(*v)
}

func main() {
	v := Vertex{3, 4}
	Print(v)

	square(&v)

	Print(v)

}
