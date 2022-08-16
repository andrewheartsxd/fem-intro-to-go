// ## Exercise 7A: Pointers

// ## Setup
// - Create a file called `exercise_7a.go` in the `07_pointers/code` directory
// - Start with the setup code snippet below

// ## Directions
// 1. Define an instance of the User struct
// 2. Write a function called `updateEmail` that takes in a `*User` type
// 3. Update the user's email to something new
// 4. Call `updateEmail()` from `main()` and verify the updated email has persisted

// ```go
// package main

// import "fmt"

// type User struct {
// 	ID                         int
// 	FirstName, LastName, Email string
// }

// // YOUR CODE HERE

// func main() {
//     fmt.Println("Pointers!")
// }
// ```

// ## Hints

// _HINT_ Check out the [go docs](https://tour.golang.org/moretypes/4) if things feel a little weird

package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func main() {
	user := User{ID: 1, FirstName: "Yo", LastName: "Momma", Email: "yo.momma@gmail.com"}
	fmt.Println(user.Email)
	updateEmail(&user, "yo.daddy@gmail.com")
	fmt.Println(user.Email)
}

func updateEmail(user *User, email string) *User {
	user.Email = email
	return user
}
