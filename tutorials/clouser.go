package main

import "fmt"

func count() {
	i := 0
	anony := func() {
		i = 0
		i += 1
		fmt.Printf("value of count :%v\n", i)
	}
	anony()
}

func main() {
	count()
	count()
}
