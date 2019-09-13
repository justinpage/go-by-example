package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p) // prints instance of struct

	fmt.Printf("%+v\n", p) // includes the struct's field names

	fmt.Printf("%#v\n", p) // syntax representation of the value

	fmt.Printf("%T\n", p) // type of a value

	fmt.Printf("%t\n", true) // boolean

	fmt.Printf("%d\n", 123) // standard base-10 formatting

	fmt.Printf("%b\n", 14) // binary representation

	fmt.Printf("%c\n", 33) // prints the corresponding character

	fmt.Printf("%x\n", 456) // provides hex encoding

	fmt.Printf("%f\n", 78.9) // basic decimal formatting

	fmt.Printf("%e\n", 123400000.0) // scientific notation
	fmt.Printf("%E\n", 123400000.0) // scientific notation

	fmt.Printf("%s\n", "\"string\"") // basic string printing

	fmt.Printf("%q\n", "\"string\"") // double quote strings

	fmt.Printf("%x\n", "hex this") // renders the string in base-16

	fmt.Printf("%p\n", &p) // prints the representation of a printer

	fmt.Printf("|%6d|%6d|\n", 12, 345) // right justified length

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // right justified length

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // include decimal precision

	fmt.Printf("|%6s|%6s|\n", "foo", "b") // right justified string

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // left justified string

	s := fmt.Sprintf("a %s", "string") // format and return
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // use io.Writer interface
}
