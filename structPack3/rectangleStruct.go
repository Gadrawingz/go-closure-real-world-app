package main

import "fmt"

func main() {

	// 1. Declare a struct
	type Rectangle struct {
		length  int
		breadth int
	}

	// 2. Declare instance rect1 and defining the struct
	rect1 := Rectangle{220, 40}

	// 3. Accessing the length and breadth of the struct
	// We use the . (dot) symbol to access the property of a struct.
	fmt.Println("LENGTH:", rect1.length)
	fmt.Println("LENGTH:", rect1.breadth)

	// 4. Calculate the area
	area := rect1.length * rect1.breadth
	fmt.Println("AREA IS: ", area)
}
