package main

import "fmt"

func main() {
	var id int = 10
	{
		id := 20 // for Shrt var declaration scope wont reflect the gloable scope
		// check id = 20
		fmt.Printf("check scope of id: %v\n", id)
		{
			id = 30
			fmt.Printf("check scope of id: %v\n", id)
		}
	}
	fmt.Printf("check scope of id: %v\n", id)
}

//o/p : shortVarDecla --- 20 and 10
