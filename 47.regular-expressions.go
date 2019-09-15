package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") // notice there is a capture group

	fmt.Println(r.MatchString("peach")) // compile optimized RegExp struct

	fmt.Println(r.FindString("peach punch")) // returns matched string

	fmt.Println(r.FindStringIndex("peach punch")) // start and end index

	fmt.Println(r.FindStringSubmatch("peach punch")) // matches and submatches

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // indexes of above

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // find all matches

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1)) // indexes of all matches and submatches

	fmt.Println(r.FindAllString("peach punch pinch", 2)) // find two matches

	fmt.Println(r.Match([]byte("peach"))) // use []byte arguments, not strings

	r = regexp.MustCompile("p([a-z]+)ch") // useful when working with constants
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // replace subsets

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
