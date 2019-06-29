package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:\t\t\t", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:\t\t\t", s)
	fmt.Println("get:\t\t\t", s[2])

	fmt.Println("length:\t\t\t", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("append:\t\t\t", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:\t\t\t", c)

	l := s[2:5]
	fmt.Println("slice one:\t\t", l)

	l = s[:5]
	fmt.Println("slice two:\t\t", l)

	l = s[2:]
	fmt.Println("slice three:\t\t", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare and initialize:\t", t)

	twoDimensions := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDimensions[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimensions[i][j] = i + j
		}
	}

	fmt.Println("two dimensions:\t\t", twoDimensions)
}
