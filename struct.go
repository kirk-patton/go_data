package main

import "fmt"

//import "math"

type Person struct {
	firstName, lastName, street string
}

type Android struct {
	Person
	Model string
}

func (p *Person) address() string {
	return fmt.Sprintf("First: %v\nLast: %v\nAddress: %v", p.firstName, p.lastName, p.street)
}

func (a *Android) model() string {
	return a.Model
}

func main() {
	person := Person{firstName: "Joe", lastName: "BudDude", street: "1234 Silly St."}
	fmt.Println(person.address())

	robot := new(Android)
	robot.lastName = "StarTrek"
	robot.firstName = "Data"
	robot.Model = "123456789"
	fmt.Println(robot.address())
	fmt.Println(robot.model())
}
