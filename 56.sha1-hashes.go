package main

import (
	"crypto/sha1"
	"crypto/md5"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	m := "md5 this string"

	h = md5.New()

	h.Write([]byte(m))

	ms := h.Sum(nil)

	fmt.Println(m)
	fmt.Printf("%x\n", ms)
}

