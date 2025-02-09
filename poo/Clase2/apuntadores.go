package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	g := 25
	fmt.Println(g)

	h := &g
	fmt.Println(h)  // muestra la direccion de memoria
	fmt.Println(*h) // muestra el valor de la variable apuntada
}
