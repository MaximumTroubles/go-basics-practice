package main

import "fmt"

func main() {
	// its how to init channel
	// channels provides tools for communation between goroutines.
	// var intCh chan int+

	// non-buffered channels init
	// var intCh chan int = make(chan int)

	// short init nb channel
	intCh := make(chan int)


	// buffered channels init 
	// it is a channel with a sized capacity
	intCh2 := make(chan int, 3)
	intCh2 <- 2
	intCh2 <- 3
	intCh2 <- 4
	// exemple of overload buffered channel (deadlock) and function main will wait until spot in chennel be free
	// intCh2 <- 5

	fmt.Println(<-intCh2)
	fmt.Println(<-intCh2)
	fmt.Println(<-intCh2)

	// We have Unidirectional channels. That can only recieve data or send data
	//Reciever only
	// var inChannel chan<- int

	//Dispatcher only
	// var outChannel <-chan int


	// We have returning type channel
	fmt.Println("Start returning type exemple")
	fmt.Println(<-createChan(22))
	fmt.Println("End returning type exemple")


	// Closing channel
	fmt.Println("Start close channel exemole")
	intChannel := make(chan int, 3)
	intChannel <- 10
	intChannel <- 3
	close(intChannel)
	// here if we try to put some we will get panic error
	// intChannel <- 5
	fmt.Println(<-intChannel)
	fmt.Println(<-intChannel)
	fmt.Println(<-intChannel)

	for i := 0; i < cap(intChannel); i++ { 
		// Here we can verify channel status. channels returns 2 value. first data, second bool value. if open = true, close = false 
		if val, opened := <-intChannel; opened { 
		   fmt.Println(val) 
		} else { 
		   fmt.Println("Channel closed!") 
		} 
  }

	fmt.Println("End close channel exemole")





	// here we are init routine and channel. Routine blocked until another routine not recieve it
	go func ()  {
		fmt.Println("Go routine starts")
		intCh <- 5
	}()
	
	// here we recieve routine by main function which is routine as well
	fmt.Println(<- intCh)
	fmt.Println("The end")
}

func createChan(n int) chan int {
	ch := make(chan int)
	go func ()  {
		ch <- n
	}()

	return ch
}