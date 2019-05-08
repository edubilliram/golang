package main

import "fmt"

func getPrintMsg() func(string) { // clouser function -  as it updates the "text" value for every iteration
	var text string = "modified"
	return func(msg string) { //anonymous functio
		text += msg
		fmt.Println(text)
	}
}

func main() {
	Print := getPrintMsg()
	Print("golang")
	Print("Gopher")

}
