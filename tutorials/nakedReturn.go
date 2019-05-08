package main

import "fmt"

func div(num, dem float32) (res float32, err error) {
	if dem == 0 {
		//res, err := 0, fmt.Errorf("res :Not nunber")
		res = 0
		err = fmt.Errorf("Not a number\n")
		return
	} else {
		res = num / dem
		err = nil
		return
	}
}

func main() {

	res, err := div(10, 5)
	if err != nil {
		fmt.Errorf("error exiting =%v\n", err)
	} else {
		fmt.Printf("res :  %v\n", res)
	}
	res1, err1 := div(10, 0)
	if err != nil {
		fmt.Errorf("error exiting =%v\n", err1)
	} else {
		fmt.Printf("res :  %v\n", res1)
	}

}
