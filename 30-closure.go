package main

import "fmt"

/*
* Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
* Harap gunakan fiture closure ini dengan bijak saat kita membuat aplikasi
**/

func main() {

	name := "Ary"
	counter := 0
	fmt.Println("Counter Awal", counter)
	increment := func() {
		fmt.Println("Increment")
		counter++ //* counter ini akan meng-increment variable counter di luar block function increment👌
	}

	changeName := func() {
		name := "Budi"
		fmt.Println("Change Name 😵")
		fmt.Println("Inner Name", name)
	}

	increment()
	increment()
	changeName()
	fmt.Println("Counter", counter)
	fmt.Println("Name", name)
}