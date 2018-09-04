package main

import (
	"errors"
	"fmt"
)

func F1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type ArgError struct {
	Arg  int
	Prob string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.Prob)
}

func F2(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := F1(i); e != nil {
			fmt.Println("F1 failed: ", e)
		} else {
			fmt.Println("F1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := F2(i); e != nil {
			fmt.Println("F2 failed: ", e)
		} else {
			fmt.Println("F2 worked: ", r)
		}
	}

	_, e := F2(42)
	if ae, ok := e.(*ArgError); ok {
		fmt.Println(ae.Arg)
		fmt.Println(ae.Prob)
	}
}
