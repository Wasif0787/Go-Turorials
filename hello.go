package main

import (
	"fmt"
	"goTutorial/testpackage"
)

func main() {
	// Integer data type declaration and initialization
	var step int = 2
	// Call to MyFunction from the testpackage with 'step' as a parameter
	testpackage.MyFunction(step)

	// Function call with parameters and return value handling
	sum := testpackage.Sum(1, 2)
	// Output the result of the Sum function
	fmt.Println("Sum is: ", sum)

	// String data type declaration and concatenation
	var firstName string = "Wasif"
	var lastName string = "Hussain"
	// Output the concatenated string of first and last name
	fmt.Println(firstName + " " + lastName)

	// Boolean data type declaration and initialization
	var isHappy bool = true
	// Output the value of the boolean variable
	fmt.Println("Am I happy?", isHappy)

	// Float data type declaration and initialization
	var pi float32 = 3.14
	// Output the value of the float variable
	fmt.Println("Value of pi is: ", pi)

	// Complex data type declaration and initialization
	var complexNumber complex64 = 1 + 2i
	// Output the value of the complex number
	fmt.Println("Value of complexNumber is: ", complexNumber)

	// Call to MyLoopFunction from the testpackage
	testpackage.MyLoopFunction()
}