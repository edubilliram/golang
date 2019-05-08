package main

import "fmt"
import "errors"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// div-- return val, err
// div(2/0) -- return 0, "Not a number"
// div(n/d) -- return val, nil
func div(num, den int) (val int, err error) {
	if den == 0 {
		err = errors.New("Not a number")
	} else {
		val = num / den
	}
	return val, err
}

func main() {
	fmt.Println(split(17))
	//if err is not nil, then print error
	//otherwise print the value
	if res, err := div(10, 0); err != nil {
		fmt.Printf("error in div: %v\n", err)
	} else {
		fmt.Printf("res in div: %v\n", res)
	}
	//fmt.Println(div(3,0))
}
