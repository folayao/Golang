package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	//Herencia
	Person
	Employee
	//Es de forma anónima

	vacation bool
}

type TemporalyTimeEmployee struct {
	//Herencia
	Person
	Employee
	//Es de forma anónima
	taxRate int
}

//interface
type PrintInfo interface {
	getMessage() string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

func NewFullTimeEmployee(id int, name string, vacation bool) *FullTimeEmployee {
	return &FullTimeEmployee{}
}
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

//principio de composición sobre herencia
func main() {
	e := FullTimeEmployee{}
	e.age = 1
	e.name = "Felipe"
	e.id = 22
	e.vacation = true
	fmt.Printf("%v", e)

	tEmployee := TemporalyTimeEmployee{}
	getMessage(tEmployee)
	getMessage(e)
	// En go no existe la herencia directa, por decir que yo heredo directamente las funciones de x clase
	// para implementar esto se necesita el uso de interfaces
	// Estas son un tipo de contratos ( Comprometerse a implementar lo definido por la interfaz)
	// que se tienen con las clases a la hora de ser implementadas
	// Go no usa una forma explicita de usar estas implementaciones
}
