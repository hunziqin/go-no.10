package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Mercedes struct {
	a int
}

type Car struct {
	Engine
	wheelCount string
	Mercedes
}

func (c *Car) numberOfWheels() {
	c.Start()
	c.Stop()
}

func (m *Mercedes) sayHiToMerkel() int {
	return m.a
}

func main() {
	M := new(Car)
	M.a = 10
	fmt.Println(M.sayHiToMerkel())
}
