package main

import "fmt"

func main(){
	// Constant adalah konstanta yang mana tidak bisa diganti value nya setelah deklarasi
	const phi = 3.14;

	// Deklarasi Multiple Constant
	const (
		firstName string = "Ary"
		lastName = "Setiyawan"
	);

	fmt.Println(phi, firstName, lastName);
}