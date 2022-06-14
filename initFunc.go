package main

import "fmt"

var name string

// function that stats before main() always.
// point of entry
func init() {
	name = "Max"
}

func main() {
	fmt.Println(name)
}
