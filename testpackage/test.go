package testpackage

import "fmt"

// MyFunction prints a message based on the provided 'step' value.
// It handles four cases:
// - If step == 1, it prints "Step 1"
// - If step == 2, it prints "Step 2"
// - If step == 3, it prints "Step 3"
// - For any other value, it prints "Invalid step"
//
// Parameters:
//   step - An integer that determines which message is printed.
//
// This function also calls a private function `myPrivateFunction`
// which is only accessible within the same package.
func MyFunction(step int) {
    fmt.Println("Welcome to MyFunction")
    
    // Conditional logic to print messages based on the 'step' value
    if step == 1 {
        fmt.Println("Step 1")
    } else if step == 2 {
        fmt.Println("Step 2")
    } else if step == 3 {
        fmt.Println("Step 3")
    } else {
        fmt.Println("Invalid step") // Handles values other than 1, 2, or 3
    }
    
    // Call to a private function within the same package
    myPrivateFunction()
}

// Sum computes and returns the sum of two integers.
//
// Parameters:
//   a - The first integer to be added.
//   b - The second integer to be added.
//
// Returns:
//   The sum of the integers 'a' and 'b'.
func Sum(a int, b int) int {
    return a + b
}

// myPrivateFunction is a private helper function that prints a message.
// It is only accessible within the `testpackage` package.
//
// This function is used internally within the package to demonstrate
// private function usage.
func myPrivateFunction() {
    fmt.Println("This is a private function")
}

// MyLoopFunction demonstrates the use of a for loop with control flow statements.
// It loops from 0 to 4, skipping the iteration when i == 2 and breaking
// the loop when i == 4.
//
// The loop prints the value of 'i' for each valid iteration, excluding
// the skipped and terminated iterations.
func MyLoopFunction() {
    // Loop from 0 to 4
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue // Skip the current iteration when i == 2
        }
        if i == 4 {
            break // Exit the loop when i == 4
        }
        // Print the value of 'i' for each valid iteration
        fmt.Println(i)
    }
}