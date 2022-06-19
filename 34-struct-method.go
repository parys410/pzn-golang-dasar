package main

import (
	"fmt"
)

/**
** Struct Method
** Struct adalah tipe data seperti tipe data lainnya, sehingga dia bisa digunakan sebagai parameter untuk function
** Kita juga bisa menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function (sebenarnya bukan karena pendefinisiannya di luar struct)
** Method adalah function seperti pada class yang juga memiliki method
 */

type Customers struct {
	Name, Address string
	Age int
	Married bool
}

func (customer Customers) sayHello() {
	fmt.Printf("Hello, My name is %s \n", customer.Name)
}

func main() {
	pelanggan1 := Customers {
		Name: "Ary",
		Address: "Abianbase",
		Age: 29,
		Married: false,
	}

	pelanggan1.sayHello()
}