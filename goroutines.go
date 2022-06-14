package main

import "fmt"

// main function it is already a goroutine with one thread, we do not see any result of factorial() method
// because main and factorial began to work in parallel. But main goroutine is main and if main is over all proccessed should be down

func main() {
	for i := 1; i < 7; i++ {
		go factorial(i)
	}

	// we add new function that expect user to press key
	int, _ := fmt.Scanln()
	fmt.Println("The End", int)
}

func factorial(n int) {
	if (n < 1) {
		fmt.Println("invid input number")
		return
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(n, "-", result)
}