package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])

	fmt.Println("length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare and initialize:", b)
	fmt.Printf("%T\n", b)

	var twoDimensions [2][3]int
	for i, x := 0, len(twoDimensions); i < x; i++ {
		for j, y := 0, len(twoDimensions[0]); j < y; j++ {
			twoDimensions[i][j] = i + j
		}
	}

	fmt.Println("twoDimensions:", twoDimensions)
}
