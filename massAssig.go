package main

import "fmt"

//global variables
var a, b, c int

func main() {
	// multiple assignment
	a, b, c := 1, 2, 3
	// flow goes in <- that direction like a,b = 2,1
	a, b = b, a
	// assignment to the void
	a, _, c = 1, 5, 3

	fmt.Println(a, b, c)
	output()
}

func output() {
	fmt.Println(a, b, c)
}
