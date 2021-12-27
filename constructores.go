package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v", e)
	// 2
	e2 := Employee{
		id:   123,
		name: "Felipe Version 2",
	}
	fmt.Printf("%v", e2)
	// 3 equivalente a la forma numero 1
	e3 := new(Employee) //  Cuando se usa new retorna un apuntador!
	e3.id = 1
	e3.name = "FELIPE e3"
	fmt.Printf("%v", *e3)
	e4 := NewEmployee(2, "Felipe", false)
	fmt.Printf("%v", *e4)

}
