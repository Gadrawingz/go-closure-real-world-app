package main

import "fmt"

func main() {

	// 1. Declare Employee struct:
	type Employee struct {
		id     int
		name   string
		salary float64
	}

	// 2. Create Instances
	// We can assign values to the fields of a struct using curly bracket and
	// assign it to a variable. This is called instance.
	// You can create multiple instances of the struct with different values.

	// 3. Struct instance
	var emp1 = Employee{1, "Joyner Lucas", 4000.75}
	var emp2 = Employee{2, "J. Cole", 8300.50}

	// 4. Printing
	fmt.Println(emp1) // {1 Joyner Lucas 4000.75}
	fmt.Println(emp2) // {2 J.Cole 8300.5}

	// 5. Declare an instance and assign values later on using a dot
	var emp3 Employee
	emp3.id = 3
	emp3.name = "Benny The Butcher"
	emp3.salary = 75650.60
	fmt.Println(emp3) // {3 Benny The Butcher 75650.6}

	// 6. EXCEPTIONS
	// NB1:
	var emp4 Employee
	emp4.salary = 84700.504
	fmt.Println(emp4) // {0  84700.504}

	// NB2:
	var emp5 Employee
	emp5.name = "Bob Marley"
	fmt.Println(emp5) // {0 Bob Marley 0}

	// 7. Struct Instantiation using Shorthand Syntax
	emp6 := Employee{23, "Dave Santana", 40000.65} // struct literal
	fmt.Println(emp6)                              // {23 Dave Santana 40000.65}

	// 7.NB: We can assign values in different order
	emp7 := Employee{id: 25, name: "Eminem"}
	fmt.Println(emp7) // {25 Eminem 0}

	emp8 := Employee{salary: 24500.55, name: "Bruno Mars"}
	fmt.Println(emp8) // {0 Bruno Mars 24500.55}

}
