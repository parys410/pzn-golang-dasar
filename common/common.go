package common

import "github.com/parys410/GoLang-Dasar/helper"

func SetAge(person *helper.Person) {
	person.Age = getAge(person.BdayYear)
}

func getAge(bdayYear int) int {
	return 2022 - bdayYear
}