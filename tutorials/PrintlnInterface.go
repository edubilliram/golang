package main

import (
	"errors"
	"fmt"
	"time"
)

type myError struct {
	time time.Time
	err  error
}

func (e myError) Error() string {
	return fmt.Sprintf("time:(%v) error: %v", e.time, e.err)
}

type person struct {
	name string
	age  int
}

func (p person) stringer() string {
	return fmt.Sprintf("Name: %v age(%v)", p.name, p.age)
}

func oops() error {
	return myError{
		time.Date(2018, 3, 28, 4, 35, 50, 1344, time.UTC),
		errors.New("let's try"),
	}
}
func main() {
	a := &person{"ram", 28}
	b := &person{"raj", 29}
	fmt.Printf("a: %v, b: %v\n", *a, *b)

	err := oops()
	if err != nil {
		fmt.Println(err)
	}
}
