package main

import "fmt"

func main() {

	// 1. Declaring a new struct
	type Student struct {
		firstName   string
		lastName    string
		indexNumber int
	}

	// 2. Assign value 2struct while creating an instance
	student1 := Student{"Daniel", "Benjamin", 30}
	fmt.Println(student1)

	student2 := Student{"Fidel", "Castro", 40}
	fmt.Println(student2)

	// 3. Or Define instance, Then assign value to struct variables
	var student3 Student
	// But, we should merge variable declaration with assignment
	student3 = Student{
		firstName:   "Jay",
		lastName:    "Polly",
		indexNumber: 30444333,
	}
	fmt.Println(student3)

}
