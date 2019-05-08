//Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
//
//The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
//
//(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
//
//(Use uint8(intValue) to convert between types.)

//https://tour.golang.org/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 { // Pic is called callback function
	p := make([][]uint8, dy) // allocating memory for two dimensional array which stores pixals
	for x, _ := range p {
		p[x] = make([]uint8, dx)
		for y, _ := range p[x] {
			p[x][y] = uint8(x ^ y)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}

//func Pic(dx, dy int) [][]uint8 {          // using for loop without range
//	r := make([][]uint8, dy)
//	for i := 0; i < dy; i++ {
//		r[i] = make([]uint8, dx)
//		for j := 0; j < dx; j++ {
//			r[i][j] = uint8((i + j) / 2)
//		}
//	}
//	return r
//}
