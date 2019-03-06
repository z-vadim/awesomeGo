package main

import (
	"fmt"
)

//Class
type Department struct {
	teamList Team
}

type Team struct {
	teamId   int
	position string
	name     string
}

type Employee struct {
	//	Department
	firstName  string
	secondName string
	salary     int
	experience int
	position   string
	teamId     int
}

func (e *Employee) getSalary() int {
	if e.experience > 2 {
		e.salary = e.salary + 200
	}
	if e.experience > 5 {
		//e.salary = (e.salary * 1.2) + 500
		e.salary = e.salary + 500
	}

	return e.salary
}

//func (e *Employee) addToTeam() {
//	e.addToTeam(e.teamId, e.position, e.firstName+e.secondName})
//}

func main() {

	///
	firstEmp := Employee{
		"Ivan",
		"Ivanov",
		200,
		3,
		"dev",
		11}

	fmt.Println(firstEmp.getSalary())
}
