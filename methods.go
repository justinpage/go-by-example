package main

import "fmt"

type Rect struct {
	Width, Height int
}

func (r *Rect) Area() int {
	return r.Width * r.Height
}

func (r *Rect) Perim() int {
	return 2*r.Width + 2*r.Height
}

func main() {
	r := Rect{Width: 10, Height: 5}

	fmt.Println("area: ", r.Area())
	fmt.Println("perim: ", r.Perim())

	rp := &r
	fmt.Println("area: ", rp.Area())
	fmt.Println("perim: ", rp.Perim())
}
