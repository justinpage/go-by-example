package main

import "fmt"
import "reflect"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	fmt.Printf("%T\n", nums)
	fmt.Printf("slice? %v\n", reflect.TypeOf(nums).Kind())
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	fmt.Printf("%T\n", nums)
	fmt.Printf("slice? %v\n", reflect.TypeOf(nums).Kind())

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array?\n", reflect.TypeOf(b).Kind())

	sum(nums...)
}
