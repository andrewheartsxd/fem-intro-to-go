// # Exercise 4a: Functions

// ## Goals

// - Practice writing functions

// ## Setup

// 1. Create a file called `exercise_4a.go` in the `04_complex_structures/code` directory.

// 2. Create a function called `average` that takes three arguments separate arguments, all of which are floats.

// 3. The function should return the average of the three arguments as a float.

// 4. Make sure to call this function from within `main()`

// # Exercise 4b: Refactoring Functions

// ## Goals

// - Practice variadic functions
// - Practice and review managing types

// ## Setup

// - Feel free to refactor the code from `exercise_4a.go`, or create a new file called `exercise_4b.go` in the same directory.

// ## Directions

// 1. Refactor your code to use a variadic function that takes in an unknown number of arguments.

// ## Hints

// _HINT_: To find the length of a collection, use `len(someCollection)`.

// _HINT_: Check the type of `len(someCollection)`

package main

import "fmt"

func average(num1 float64, num2 float64, num3 float64) float64 {
	return (num1 + num2 + num3) / 3
}

func averageVariadic(nums ...float64) float64 {
	length := len(nums)
	var sum float64
	for _, value := range nums {
		sum += value
	}
	return sum / float64(length)
}

func main() {
	fmt.Println(average(10, 20, 30))
	fmt.Println(averageVariadic(10, 20, 30))
}
