package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

// Group represents a set of userss
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

	fmt.Println(u)

	fmt.Println(describeUser(u))

	g := Group{role: "admin", users: []User{u, u2}, newestUser: u2, spaceAvailable: true}
	fmt.Println(describeGroup(&g))
	fmt.Println(g)
}

func describeGroup(g *Group) string {
	if len(g.users) >= 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// - Let's say that in `describeGroup`, we only want to accept new users (represented by the `spaceAvailable` field) if there are fewer than 2 existing users in the group.
// ## Directions
// 1. In the `describeGroup`, update `spaceAvailable` to be `false` if there are 2 or more users.
// 2. Reprint the variable `g` at the end of `main()`. What do you notice?
