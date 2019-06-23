package main

import "fmt"

func main() {
	// Basic
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Break
	for {
		fmt.Println("loop")
		break
	}

	// Continue 
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
