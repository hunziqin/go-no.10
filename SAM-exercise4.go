package main

import "fmt"

type A struct {
	int
}

type B struct {
	string
}

type C struct {
	flo float32
	A
	B
}

func main() {
	TT := C{1.0, A{2}, B{"hunzi"}}
	fmt.Println(TT)
	fmt.Println(TT.flo, TT.int, TT.string)
}
