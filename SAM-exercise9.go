package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float64
	c string
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}

func (tn *T) String() string {
	return strconv.Itoa(tn.a) + "/" + strconv.FormatFloat(tn.b, 'f', 2, 64) + "/" + tn.c
}
