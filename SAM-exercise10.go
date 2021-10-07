package main

type Day struct {
	d []string
}

func main() {
	day := new(Day)
	day.d = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
}

func (dn *Day) String() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
}
