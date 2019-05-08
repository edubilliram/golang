package main

import "fmt"

type Int struct{ int }

func main() {
	v := 0
	vv := &v
	*vv++
	fmt.Println("value vv: ", *vv)

	/// same thing with new()
	vv = new(int)
	*vv++
	fmt.Println("using new -->- value vv: ", *vv)

	/* /// this is not possiable /////////
	 v := &int
	 *v++
	fmt.Println("value vv: ", *vv)
	*/ ////////////

	// using "new" for struct
	r := new(Int)
	r.int = 1
	r.int++
	fmt.Println("using new for struct----> value of r: ", r)
}
