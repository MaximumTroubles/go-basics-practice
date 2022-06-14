package main

import "fmt"

func main() {
	// we are inti slice with limit of 10 and inside init another slice
	matrix := make([][]int, 10)

	counter := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter
		}
		fmt.Println(matrix[x])
	}

	makeFor()
}

func makeFor() {
	x := 0
	for x < 10 {
		x++
		fmt.Println(x)
	}
}
