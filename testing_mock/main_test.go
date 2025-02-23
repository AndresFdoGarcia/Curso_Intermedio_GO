package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "123456789",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "Developer",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Andres",
						Age:  20,
						DNI:  "123456789",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "Andres",
					Age:  20,
					DNI:  "123456789",
				},
				Employee: Employee{
					Id:       1,
					Position: "Developer",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, item := range table {
		item.mockFunc()
		ft, err := GetFullTimeEmployeeById(item.id, item.dni)

		if err != nil {
			t.Errorf("Error when getting employee: %s", err)
		}

		if ft.Name != item.expectedEmployee.Person.Name {
			t.Errorf("Error, got %s expected %s", ft.Name, item.expectedEmployee.Person.Name)
		}

		if ft.Age != item.expectedEmployee.Person.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, item.expectedEmployee.Person.Age)
		}

		if ft.Employee.Position != item.expectedEmployee.Employee.Position {
			t.Errorf("Error, got %s expected %s", ft.Employee.Position, item.expectedEmployee.Employee.Position)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
