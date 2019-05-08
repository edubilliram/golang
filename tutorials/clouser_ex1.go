package main

import "fmt"

func ApplyToArray(fun func(int) error, numbers []int) (err error) {
	for i := 0; i < len(numbers); i++ {
		err := fun(numbers[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func sum(intss []int) (sum int) {
	sum = 0
	fun := func(nums int) error {
		sum = sum + nums
		return nil
	}
	ApplyToArray(fun, intss)
	return sum

}

func main() {
	ints := []int{1, 2, 3, 4}
	res := sum(ints)
	fmt.Println(res)
}
