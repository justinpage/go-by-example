package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["keyOne"] = 7
	m["keyTwo"] = 13

	fmt.Println("map:", m)

	valueOne := m["keyOne"]
	fmt.Println("valueOne:", valueOne)

	fmt.Println("length:", len(m))

	delete(m, "keyTwo")
	fmt.Println("map:", m)

	_, present := m["keyTwo"]
	fmt.Println("present:", present)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
