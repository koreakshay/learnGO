package main

import "fmt"

// Struct is a collection of fields
type Employee struct {
	name 	 string
	age 	 int
	Address
}

type Address struct {
	street string
	city string
}

func main() {
	address := Address {
		street: "123 Main Street",
		city: "New York",
	}
	employee1 := Employee{
		name: "Akshay",
		age: 30,
		Address: address,
	}
	fmt.Println("Employee name:", employee1.name)
	fmt.Println("Employee age:", employee1.age)
	fmt.Println("Employee address:", employee1.Address.street)

	job := struct {
		title string
		salary int
	} {
		title: "SWE",
		salary: 100000,
	}
	fmt.Println(job.title, job.salary)
}