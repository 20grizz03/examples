package main

import "fmt"

func main() {
	v := 5
	p := &v

	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v // *p = &v
}
