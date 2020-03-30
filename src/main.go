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
	path, _ := os.Getwd()
	fmt.Printf("Code dir is: %s\n", path)
	s := filepath.Join("/", "usr", "bin")
	err := os.Chdir(s)
	check(err)

	c, err := ioutil.ReadDir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Option 1
	examplePath1 := "dir" + string(os.PathSeparator) + "example"
	fmt.Println("PathSeparator: " + examplePath1)

	// Option 2
	examplePath2 := filepath.FromSlash("dir/example")
	fmt.Println("FromSlash: " + examplePath2)
}
