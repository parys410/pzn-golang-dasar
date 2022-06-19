package main

import (
	"fmt"
	"sort"
)

/**
** Package sort
** adalah package yang berisikan utilitas untuk proses pengurutan
** Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
** https://golang.org/pkg/sort/
 */

type User struct {
	Name string
	Age int
}

type UserSlice []User //! Type Allias

//! Kontrak untuk Interface Sort
func (userSlice UserSlice) Len() int {
	return len(userSlice)
}
func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}
func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}


func main() {
	users := []User {
		{ "Ary", 30 },
		{ "Budi", 35 },
		{ "Joko", 25 },
		{ "Adit", 23 },
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}