package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	{
		i = 4
		var j int
		if i == 4 {
			j = 3
		}
		fmt.Println(i, j, c, python, java)
	}
	fmt.Println(i, j, c, python, java)

}
