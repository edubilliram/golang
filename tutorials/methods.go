//similar classes
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

func (v Vertex) Print() {
	fmt.Printf("vertex: {x: %v, y: %v}\n\n", v.X, v.Y)
}

func (v *Vertex) squarePointerReceiver() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
	fmt.Printf("squarePointerReceiver() ")
	v.Print()
}
func (v Vertex) squareValueReceiver() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
	fmt.Printf("squareValueReceiver() ")
	v.Print()
}

func main() {
	v := Vertex{3, 4}
	v.Print()

	v.squareValueReceiver()
	fmt.Printf("main() ")
	v.Print()

	v.squarePointerReceiver()
	fmt.Printf("main() ")
	v.Print()

	v.squareValueReceiver()
	fmt.Printf("main() ")
	v.Print()
}
