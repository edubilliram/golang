package main

import "fmt"

func main() {
	sum := 0
	nums := []int{0, 1, 2, 3, 4, 5}

	for i, num := range nums {
		num += num
		fmt.Printf("lets print %v, %v\n", num, i)
	}

	for _, num := range nums {
		sum += num
		fmt.Printf("print sum and num : %d, %d\n", sum, num)
	}
	fmt.Printf("sum:%v\n", sum)

	for i, num := range nums {
		if i%2 == 0 {
			fmt.Printf("even number : %v\n", num)
		} else {
			fmt.Printf("odd number :%v\n", num)
		}
	}

}
