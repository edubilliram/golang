package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello world")
}

func main() {

	go hello()
	//synchronizing without Channels.
	//synchronizing using time.sleep , wait for 1 sec to finish  go routine and then go to next line.
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
