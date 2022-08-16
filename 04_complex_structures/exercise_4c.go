// # Exercise 4c: Functions, Arrays, Slices, Maps,

// ## Goals

// - Practice working with arrays, slices, and maps
// - Review `range`

// ## Setup

// - Create a file called `exercise_4c.go` in the `04_complex_structures/code` directory.

// ## Directions

package main

import "fmt"

// ### Part 1
// 1. Instantiate an array of scores
//     - The array should have at least 5 elements of type `float64`
// 2. Write a function that calculates and returns the average score (also a float)
//     - Use the `range` keyword

func averageScore(scores [5]float64) float64 {
	var total float64
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))

}

//	func main() {
//		scores := [5]float64{2.3, 1.2, 5.5, 6.1, 0.6}
//		average := averageScore(scores)
//		fmt.Println(average)
//	}

// ### Part 2
// 1. Define a map that contains a set of pet names, and their corresponding animal type. i.e.: `"fido": "dog"`.
// 2. Write a function that takes a string argument and returns a boolean indicating whether or not that key exists in your map of pets.
func inMap(name string, mapToCheck map[string]string) bool {
	_, ok := mapToCheck[name]
	return ok
}

// func main() {
// 	pets := map[string]string{"fido": "dog", "socks": "cat", "boots": "cat", "barkley": "dog"}
// 	fmt.Println(inMap("fido", pets))
// 	fmt.Println(inMap("yeller", pets))
// }

// ### Part 3
// 1. Instantiate a slice that has an initial value of a collection of groceries.
// 2. Write a function that takes one or more groceries as strings and appends them to the slice, printing out the resulting list of groceries.

func addGroceries(groceries []string, newGroceries ...string) []string {
	foods := groceries
	for _, newGrocery := range newGroceries {
		foods = append(foods, newGrocery)
	}
	return foods
}

func main() {
	var groceries = []string{"milk", "eggs", "bread", "meat"}
	groceryList := addGroceries(groceries, "beans", "veg")
	fmt.Println(groceryList)
}
