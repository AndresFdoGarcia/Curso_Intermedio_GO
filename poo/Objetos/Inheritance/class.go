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
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juan"
	ftEmployee.age = 20
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
}
