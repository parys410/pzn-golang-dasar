package main

import "fmt"

/**
** Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan.
** Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
 */

func random() interface{} {
	return "ok"
}

func main() {
	result := random()
	
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}