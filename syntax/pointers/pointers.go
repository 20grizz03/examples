package main

import "fmt"

type Person struct {
	Name string
}

func changeName(p *Person) {
	p = &Person{
		Name: "John",
	}
	// создаётся новый указатель, который не будет менять значение
}

func main() {
	p := &Person{
		Name: "bob",
	}
	fmt.Println(p.Name)
	changeName(p)
	fmt.Println(p.Name)
	// обращается к старому указателю
}

// решение
func changeName2(p *Person) {
	p.Name = "John"
}
