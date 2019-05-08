package main

import "fmt"

type ipadder [4]byte

//string() is a method defined in Stringer interface, so here ipadder implements Stringer
func (i ipadder) String() string {
	return fmt.Sprintf("%v.%v.%v.%v\n", i[0], i[1], i[2], i[3])
}

//without stringer
func (ip ipadder) modString() string {
	return fmt.Sprintf("%v.%v.%v.%v\n", ip[0], ip[1], ip[2], ip[3])
}
func main() {
	m := map[string]ipadder{
		"google": {127, 27, 2, 4},
		"yahoo":  {138, 38, 3, 8},
	}

	fmt.Printf("vlaue :%v\n", m)

	for name, value := range m {
		fmt.Printf("name: %v, value: %v\n", name, value)
	}
	// without using stringer interface, we have to call the particular method and pass variables everytime whenever we use fmt.
	ip := ipadder{1, 2, 3, 4}
	host := ip.modString()
	fmt.Println(host)
}

/*
unc getStatus(jobStatus string) string {
        status := make(map[string]string)
        status = map[string]string{
                "New":        "New",
                "Dispatched": "Assigned",
                "WIP":        "In Pogress",
                "Re-assign":  "Pending",
                "Resolved":   "Resolved",
                "Closed":     "Resolved",
        }

        return status[jobStatus]

}

*/
