package main

import "fmt"

func main() {
	var csvData [][]string

	csvData = [][]string{
		[]string{"k1", "v11", "v12"},
		[]string{"k1", "v21", "v22"},
		[]string{"k2", "v21", "v22"},
		[]string{"k3", "v31", "v32"},
	}
	res := make(map[string][]string)
	for _, d := range csvData {
		key := d[0]
		vls := d[1:]
		//res[key] = vls    this direct assignment overrides key "K1" and it "vls", o/p will be := map[k1:[v21 v22] k2:[v21 v22] k3:[v31 v32]]
		// so we have to chcek the "key" in the map, and append if it has some data in it.

		if m, ok := res[key]; ok {
			res[key] = append(m, d[1], d[2]) // not preffered  append(m, vls)
		} else {
			res[key] = vls
		}

	}
	fmt.Println(res)
	//fmt.Println(csvData)

	// good way is to store v11, v12 as object, as k1 has multiple set of vlues
}
