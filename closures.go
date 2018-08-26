package main

import "fmt"

func intSeq() (increment func() int) {
	i := 0
	increment = func () int {
		i++
		return i
	}
	return
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}

