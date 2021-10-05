package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	tow1 := new(TwoInts)
	tow1.a = 12
	tow1.b = 10

	fmt.Printf("The sum is: %d\n", tow1.AddThem())
	fmt.Printf("Add them to the param: %d\n", tow1.AddToParam(20))

	tow2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", tow2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
