package main

import "fmt"

func getPrintMsg(name string) func() { /// getPrintMsg() is clouser function
	text := "modified" + name

	return func() { /// func() is anaonymous function
		fmt.Println(text)
	}
}

func main() {
	Print := getPrintMsg("golang")
	Print()
	Print()

}
