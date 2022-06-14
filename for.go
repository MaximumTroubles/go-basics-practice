package main

import "fmt"

func main() {
	fmt.Println(findMin(123, 1, 1, 2, 535, 2341, 1231, 45345, -45, 123, 234, 6575, 567, 35, 7, 57))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}
