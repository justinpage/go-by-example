package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) grow() {
	r.width = r.width * 2
	r.height = r.height * 2
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	rp.grow()
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	// same value because it is using the same underlying array
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
}
