package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("/tmp/subdir", 0755)
	check(err)

	defer os.RemoveAll("/tmp/subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("/tmp/subdir/file1")

	err = os.MkdirAll("/tmp/subdir/parent/child", 0755)
	check(err)

	createEmptyFile("/tmp/subdir/parent/file2")
	createEmptyFile("/tmp/subdir/parent/file3")
	createEmptyFile("/tmp/subdir/parent/child/file4")

	c, err := ioutil.ReadDir("/tmp/subdir/parent")
	check(err)

	fmt.Println("Listing /tmp/subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("/tmp/subdir/parent/child")
	check(err)

	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing /tmp/subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting /tmp/subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", p, info.IsDir())
	return nil
}
