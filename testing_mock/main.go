package main

import (
	"time"
)

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{
		Name: "Andres",
		Age:  20,
		DNI:  "123456789",
	}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{
		Id:       1,
		Position: "Developer",
	}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p

	return ftEmployee, nil
}
