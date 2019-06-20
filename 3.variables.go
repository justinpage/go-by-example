package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Printf("%T\n", d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Printf("%T\n", f)
}
