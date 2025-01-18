package main

import "fmt"

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func check(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	check(e1)

	var e *errorString
	check(e)

	e = &errorString{}
	check(e)

	e = nil
	check(e)
}
