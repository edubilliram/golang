package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// here just modifing the literals of the struct Vertex, but not trying to return/ or print
func (v *Vertex) Scale(f float64) *Vertex { // return Vertex
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}
func main() {
	fmt.Println("vim-go")
	v := Vertex{3, 4}
	//a := v.Scale(10)
	v.Scale(10)
	//fmt.Println(a)
	fmt.Println(v.Abs())
}
