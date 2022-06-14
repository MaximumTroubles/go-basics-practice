package main

import (
	"fmt"
	"reflect"
)

// go is statically typed, compiled programming language

func main() {
	// can't be unused variable
	// init var := short one
	message := "I will be go developer soon \n"
	// init var long one (if init imminently no need to write type. go defined it by itself)
	var text1 string
	text1 = "some text"
	var text2 = "long var init \n"
	// init constant
	const pi = 3.141592
	// rewrite var allow only with =
	message = "I will be go developer soon 2 \n"

	// 0 value variable
	// for string it is empty sting ""
	var zeroText string
	// 0 for int
	var num int
	// 0 for unsigned int => only positive number
	var uNum uint
	// 0 for float
	var floatNum float32
	// false for boolean var
	var bake bool
	// how bytes works
	var a byte = 62
	var b rune = '>'

	fmt.Printf(message)
	fmt.Printf(text1)
	fmt.Printf(text2)
	fmt.Println(pi)
	fmt.Println(reflect.TypeOf(pi))

	fmt.Println(zeroText)
	fmt.Println(num)
	fmt.Println(uNum)
	fmt.Println(floatNum)
	fmt.Println(bake)

	fmt.Printf("%c\n", a)
	fmt.Println(b)
}
