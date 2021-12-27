package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hola")

	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println("h", h)
	fmt.Println("h operador estrella nos entrega el valor de la referencia ", *h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
