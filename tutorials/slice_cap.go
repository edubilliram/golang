package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]                     // underlying array is [6]int{2,3,5,7,11,13}
	fmt.Println(s, len(s), cap(s)) // for cap always relay on UNDERLYING array end value to  modified slice intial value

	//as we're doing re-slicing
	s = s[:2] // underlying array is [5]int{3,5,7,11,13}
	fmt.Println(s, len(s), cap(s))

	s = s[1:] // underlying array is [5]int{3,5}
	// s = s[1:5] underlying array is [5]int{3,5,7,11,13}
	fmt.Println(s, len(s), cap(s))

	// s = s[0:5]  -- gives panic becuse it's cap is 4
	s = s[0:4] // underlying array is [4]int{5,7,11,13}
	fmt.Println(s, len(s), cap(s))

	s1 := s[1:3]
	fmt.Println(s1, len(s1), cap(s1))
}
