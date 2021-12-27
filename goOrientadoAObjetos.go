package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	// Creacion de una instancia de un empleado
	empl := Employee{}
	fmt.Printf("%v", empl)
	empl.id = 45678
	empl.name = "Felipe"
	fmt.Printf("%v", empl)
}
