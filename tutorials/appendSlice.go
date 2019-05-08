package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	condition := "outage"
	Time := "2018-02-02T13:24:17-05:00"
	severity := "Low"
	summary := "NES - CNY - Multi Hub - Multi Node - HAZCON - L"
	sumSlc := strings.Split(summary, "-")

	//sumSlc = sumSlc[:len(sumSlc)-2]
	sumSlc = append(sumSlc[:len(sumSlc)-2], strings.ToUpper(condition), severity[:1])

	s := strings.Join(sumSlc, " - ")

	y := s + "\n" + severity
	tm, err := time.Parse(time.RFC3339, Time)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(tm)
	fmt.Println(s)
	fmt.Println(y)

}
