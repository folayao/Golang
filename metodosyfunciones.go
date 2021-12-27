package main

import "fmt"

type Employee struct {
	id   int
	name string
}

//Reciever functions
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetId() int {
	return e.id
}

func main() {
	// Creacion de una instancia de un empleado
	empl := Employee{}
	fmt.Printf("%v", empl)
	empl.id = 45678
	empl.name = "Felipe"
	fmt.Printf("%v", empl)
	//Despues
	empl.SetId(5)
	empl.SetName("FelipeOlaya")
	fmt.Printf("%v", empl)
	fmt.Println(empl.GetId())
	fmt.Println(empl.GetName())
}
