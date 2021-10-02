package main

import "fmt"

type Vcard struct {
	name  string
	Pcode string
	date  string
}

type Address struct {
	Vcard
}

func main() {
	var vcard Vcard
	vcard.name = "hunzi"
	vcard.Pcode = "537128"
	vcard.date = "2001"
	fmt.Printf("%s %s %s\n", vcard.name, vcard.Pcode, vcard.date)
}
