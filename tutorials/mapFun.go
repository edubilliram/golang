package main

import (
	"fmt"
)

/*var status := map[string]string{
	"New":        "New",
	"Dispatched": "Assigned",
	"WIP":        "In Pogress",
	"Re-assign":  "Pending",
	"Resolved":   "Resolved",
	"Closed":     "Resolved",
}*/

func gts(jstatus string) string {
	status := make(map[string]string)
	status = map[string]string{
		"New":        "New",
		"Dispatched": "Assigned",
		"WIP":        "In Pogress",
		"Re-assign":  "Pending",
		"Resolved":   "Resolved",
		"Closed":     "Resolved",
	}

	return status[jstatus]
}

func main() {
	fmt.Println("Hello, playground")
	s := gts("Resolved")
	fmt.Println("status: ", s)

	//	m := status["New"]
}
