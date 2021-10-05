package main

import "fmt"

type employee struct {
	salary float32
}

func main() {
	s := new(employee)
	s.salary = 10
	fmt.Printf("Add is: %f\n", s.giveRaise())
}

func (sum *employee) giveRaise() float32 {
	return sum.salary + sum.salary*0.1
}
