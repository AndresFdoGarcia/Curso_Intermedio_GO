package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// func getValues(x int) (int, int , int){
// 	return 2 * x, 3 * x, 4 * x
// }

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	printNames("Juan", "Carlos", "Andres")
	fmt.Println(getValues(2))
}
