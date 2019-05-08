package main

import "fmt"

type values struct {
	a string
	b string
}

func main() {
	csvData := [][]string{
		[]string{"k1", "v11", "v12"},
		[]string{"k1", "v21", "v22"},
		[]string{"k2", "v21", "v22"},
		[]string{"k3", "v31", "v32"},
	}
	resol := make(map[string][]values)
	//var k []string
	//var v []values
	for _, c := range csvData {
		ky := c[0]
		objVal := values{a: c[1], b: c[2]}
		//resol[key] = []values{objVal} to not override we use --- append vlaues in a key
		if v, ok := resol[ky]; ok {
			resol[ky] = append(v, objVal)
		} else {
			resol[ky] = []values{objVal}
		}
	}
	/* TO READ FROM
	for _, ks := range k {
		k = append(k, ks)
		fmt.Println("k:", k)
	}
	for _, mp := range k {
		fmt.Printf("k: %v, Value: %v\n", mp, resol[mp])
	}*/
	fmt.Println(resol)
	for key, values := range resol {
		fmt.Println("keys: ", key, "values: ", values)
	}
}
