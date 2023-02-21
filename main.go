package main

import (
	"fmt"
	"strings"
)

type User struct {
	name string
	age uint8
	email string
}

func (u User) String() string {
	return fmt.Sprintf("%v (%v years). Contact info: %v", u.name, u.age, u.email)
}

func main() {
	users := []User{
		{
			name: "Mike",
			age: 32,
			email: "mike@gmail.com",
		},
		{
			name: "John",
			age: 54,
			email: "john@gmail.com",
		},
		{
			name: "Abobus",
			age: 2,
			email: "abobus@gmail.com",
		},
	}
	userInfo, ok := searchUser("mike", users)
	printSearchResult(userInfo, ok)
	userInfo, ok = searchUser("Mike", users)
	printSearchResult(userInfo, ok)
	userInfo, ok = searchUser("Abobusss", users)
	printSearchResult(userInfo, ok)
}

func searchUser(name string, users []User) (user User, ok bool) {
	name = strings.ToLower(name)
	userMap := make(map[string]User)
    for _, u := range users {
        userMap[strings.ToLower(u.name) ] = u
    }
	user, ok = userMap[name]
	return
}

func printSearchResult(info User, ok bool) {
	if ok {
		fmt.Println(info)
	} else {
		fmt.Println("User not found")
    }
}