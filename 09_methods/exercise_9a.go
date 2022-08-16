// # Exercise 9: Methods

// ## Goals
// - Practice writing a method

// ## Setup
// - Revisit the Structs lesson (`06_structs/structs.md`) and use the `06_structs/exercise_6a.md` exercise for reference.
// - Copy the code from the solution of `06_structs/exercise_6a` into a new file called `exercise_9a.go` as a starting point.

// ## Directions
// 1. Refactor the `describeUser` code to be a method (this should be a repetition of what was just completed in the course).
// 2. Modify the `describeGroup` function to be a method called `describe()` that receives a `Group` type.
// 3. Modify the main function to reflect those changes

package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func (user *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return desc
}

// Group represents a set of userss
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func (g *Group) describe() string {
	if len(g.users) >= 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

	fmt.Println(u)

	fmt.Println(u.describe())

	g := Group{role: "admin", users: []User{u, u2}, newestUser: u2, spaceAvailable: true}
	fmt.Println(g.describe())
	fmt.Println(g)
}
