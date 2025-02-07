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
	// Empleado 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	// Empleado 2
	e2 := Employee{
		id:       1,
		name:     "Juan",
		vacation: false,
	}
	fmt.Printf("%v\n", e2)
	// Empleado 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	// Se le dan valores luego de instanciado
	e3.id = 3
	e3.name = "Pedro"
	e3.vacation = true
	fmt.Printf("%v\n", *e3)
	// Empleado 4
	e4 := NewEmployee(4, "Andres", false)
	fmt.Printf("%v\n", *e4)
}
