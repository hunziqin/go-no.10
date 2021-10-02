package main

import "fmt"

type Rectangle struct {
	a int
	b int
}

func (area *Rectangle) Area() int {
	return area.a * area.b
}

func (area *Rectangle) Perimeter() int {
	return area.a*2 + area.b*2
}

func main() {
	area := Rectangle{a: 1, b: 2}
	fmt.Println("Area:\n", area.Area())
	fmt.Println("Perimeter:\n", area.Perimeter())
}
