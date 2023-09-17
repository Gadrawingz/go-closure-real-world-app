package main

import "fmt"

/*
STRUCTURE
---------
=> Struct(structure) is used to store variables of different data types.
=> It's used 2store multiple values of different data types in a single variable.

=> Go does not support class unlike other object-oriented languages
like C# and Java. So, the struct is used like the class in go.

=> Struct in Go is declared using the type and struct keywords.
=> Blessed by God through Gad Iradufasha

// StructName is the name of the structure
type StructName struct { // struct keyword is used to define a structure
	// structure definition
	<fieldName1> dataType
	<fieldName2> dataType
	<fieldName3> dataType
}
======================================================
*/

func main() {

	/*1. CREATING A SHIT
	--------------------
	- Suppose we want to store the name and age of a person.
	- We can create two variables: name and age and store value.
	- However, suppose we want to store the same information of multiple people.
	- In this case, creating variables for a person might be a tedious task.
	- We can create a struct that stores the name and age to overcome this.
	- There4, we can use this same struct for every person
	*/

	type Person struct {
		// structure definition, w/c contains 2 variables
		name string
		age  int
	}

	// 2. Creating a Struct instance
	// =============================
	// A struct definition is just a blueprint. and To use a struct,
	// we need to create an instance of it like this:
	var person1 Person

	// 3. Using the person1 instance to access&define the struct properties.
	// Defining the value of name & age
	person1 = Person{"Gad", 10}
	fmt.Println(person1) // {Gad 10}

	person1 = Person{"Bruno", 34}
	fmt.Println(person1) // {Bruno 34}
}
