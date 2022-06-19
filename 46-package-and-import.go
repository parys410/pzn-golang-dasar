package main

/**
** Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama
** Jika kita ingin mengakses file Go-Lang yang berada di luar package, maka kita bisa menggunakan import
 */

import (
	"github.com/parys410/GoLang-Dasar/helper"
)

func main() {
	helper.SayHelloInHelper("Ary")
}