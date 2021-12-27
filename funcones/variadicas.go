package main

import "fmt"

func suma(a, b int) int {
	return a + b
}
func sumaVariadica(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printListNames(names ...string) {
	for _, name := range names {
		fmt.Println(name + ", ")
	}
}

func getValues(i int) (int, int, int) {
	return i * 2, i * 3, i * 4
}

func getValuesReturnedNames(i int) (doble int, triple int, quadro int) {
	doble = i * 2
	triple = i * 3
	quadro = i * 4
	return
}

func main() {
	// funciones variadicas y retornos con nombre
	fmt.Println(suma(1, 2))
	// En caso de que see agregar otro numero a la funcion será imposible si no es una funcion variadica
	fmt.Println(sumaVariadica(1), "Suma variadica")
	fmt.Println(sumaVariadica(1, 2), "Suma variadica")
	fmt.Println(sumaVariadica(1, 3, 4), "Suma variadica")
	fmt.Println(sumaVariadica(1, 2, 3, 4, 5), "Suma variadica")
	printListNames("Felipe", "Daniel", "Santiago")
	getValues(1)
	fmt.Println(getValuesReturnedNames(1))
}
