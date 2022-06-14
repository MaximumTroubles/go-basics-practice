package main

import "fmt"

func main() {
	//mapFunc("Alex")
	defer handlePanic()
	mapInit()
}

func mapFunc(name string) {
	users := map[string]int{
		"Max":   23,
		"Vasil": 13,
		"Petro": 33,
	}

	//adding to map
	users["Volodimir"] = 45

	//cycle for
	for key, value := range users {
		fmt.Println("From for cycle ", key, value)
	}

	//delete element from map
	delete(users, "Petro")

	age, exists := users[name]
	if exists {
		fmt.Println(name, age)
	} else {
		fmt.Println("user doesnt exist")
	}

	fmt.Println("How map looks like: ", users)
	fmt.Println("Call value by its key: ", users["Max"])
	fmt.Println("Call not existed key: ", users["Ivan"]) // for int 0
}

func mapInit() {
	// if we create map without init like this, we can't put element in, panic provided

	//var users map[string]int
	//users["ivan"] = 15

	// this is a right way how to init empty map
	users2 := make(map[string]int)
	users2["Max"] = 27

	fmt.Println(users2)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic happened")
	}
}
