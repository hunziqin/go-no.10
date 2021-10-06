package main

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(s string) {
	b.id += s
}

func (b *Base) String() string {
	return b.id
}

type Person struct {
	Base
	FirstName, LastName string
}

type Employee struct {
	Person
	salary string
}

func (e *Employee) employee() string {
	return e.id
}

func main() {
	b1 := new(Employee)
	b1.id = "hunzi"
	b1.SetId("qin")
	fmt.Println(b1.employee())
}
