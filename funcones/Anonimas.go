package main

// Funciones Anónimas CLASE 13
import "fmt"

func main() {
	x := 5
	//fUNCION Anónima
	y := func() int {
		return x * 2
	}() // Estos parentesis llaman a la funcion por si misma

	fmt.Println(y)

	// funciones variadicas y retornos con nombre

}
