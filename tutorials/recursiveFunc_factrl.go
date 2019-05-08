package main

import "fmt"

func fac(x int) (res int, err error) {
	if x < 0 {
		return 0, fmt.Errorf("invalid input\n")
	}

	if x == 0 {
		return 1, nil
	}
	// we want to assign res(return value) and we don't care about err-- > so go with "_"
	temp, _ := fac(x - 1)
	res = x * temp
	return res, nil
}

func main() {

	if res, err := fac(4); err != nil {
		fmt.Printf("invalid input : %v\n", err)

	} else {
		fmt.Printf("factorial of x is %v", res)
	}
}
