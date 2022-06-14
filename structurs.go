package main

import "fmt"

func main() {
	// struct is custom type that can have properties and methods. Based on struct we can create objects. Similar to Class in PHP

	// simple way to create and init struct and output example
	//user := struct {
	//	name string
	//	age  int
	//	sex  string
	//}{"Max", 27, "Male"}

	// create new object User, and we can copy it
	//user := User{"Max", 27, "Male"}
	//user2 := User{"Ivan", 37, "Male"}

	// create with constructor
	//user3 := NewUser("Vasyl", "Male", 20)

	// call specific properties
	//fmt.Println(user.Name)

	// call struct method
	//user3.printUserInfo("Max")

	// example how we do it PHP with Entities. Create empty and then fill it up with values
	user4 := User{}
	user4.setName("Max")
	user4.setAge(13)

	//fmt.Println(user, user2)   // display only values
	//fmt.Printf("%+v\n", user)  // extended struct print
	//fmt.Printf("%+v\n", user2) // extended struct print
	//fmt.Printf("%+v\n", user3) // extended struct print
	fmt.Printf("%+v\n", user4.getName()) // extended struct print
	fmt.Printf("%+v\n", user4.getAge())  // extended struct print
	fmt.Printf("%+v\n", user4)           // extended struct print

	if user4.Age.isAdult() {
		fmt.Println("Look how u grow up")
	} else {
		fmt.Println("Childhood it is prefect time")
	}
}

// Age We also can create our own type: Age
type Age int

// For own type provide own additional methods
func (a Age) isAdult() bool {
	return a >= 18
}

// User often way to create struct
type User struct {
	Name string
	Age  Age
	Sex  string
}

// NewUser User constructor
func NewUser(name, sex string, age int) User {
	return User{
		Name: name,
		Age:  Age(age),
		Sex:  sex,
	}
}

/**
we have 2 type of receiver (value receiver u User) and (pointer receiver u *User) when we call it,
inside we are working with link to object (pointer), so we decided to change name, so its return new value and
will continue like that, we do not copy object we provide pointer to it. Address in RAM memory cell
**/
func (u *User) printUserInfo(name string) {
	u.Name = name
	fmt.Println(u.Name, u.Sex, u.Age)
}

// example of getters and setters
func (u User) getName() string {
	return u.Name
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) setAge(age int) {
	u.Age = Age(age)
}

func (u User) getAge() Age {
	return u.Age
}

// DumbDatabase Example of constructor usage
// We have "map" in our Class, but we can't init it in the future, then we crate constructor and init map with make()
// method and return Class by link (not copy)
type DumbDatabase struct {
	m map[int]string
}

func NewDumpDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[int]string),
	}
}
