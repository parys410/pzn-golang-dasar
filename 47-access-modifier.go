package main

/**
** Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
** Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
** Jika namanya diawali dengan huruf besar, maka artinya bisa diakses dari package lain. Jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain
 */

import (
	"github.com/parys410/GoLang-Dasar/common"
	"github.com/parys410/GoLang-Dasar/helper"
)

func main() {
	newPerson := helper.Person {
		Name: "Ary Setiyawan",
		BdayYear: 1992,
	}
	common.SetAge(&newPerson)
	newPerson.Greetings()
}