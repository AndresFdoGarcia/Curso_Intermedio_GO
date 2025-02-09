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
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Andres"
	ftEmployee.age = 20
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee)
	fmt.Println("--------------------------------")
	//Empleado termporal
	tEmployee := TemporaryEmployee{}
	tEmployee.name = "Carlos"
	tEmployee.age = 20
	tEmployee.id = 1
	tEmployee.taxRate = 10
	fmt.Printf("%v\n", tEmployee)
	GetMessage(tEmployee)
}
