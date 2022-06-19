package main

import "fmt";

func main(){
	
	// var (namaVariable) (typeVariable)
	var name string;
	var fullName = "Putu Ary Setiyawan";
	var age int8 = 30;

	// Kata kunci var tidak wajib apabila mengubah = menjadi := (saat pertama x deklarasi)
	fullNameNew := "Putu Ary Setiyawan";

	// Membuat banyak variable sekalian
	var (
		firstName = "Eko Kurniawan"
		lastName = "Khannedy"
	);

	name = "Ary Setiyawan";
	fmt.Println(name);

	name = "Eko khannedy";
	fmt.Println(name);
	fmt.Println(fullName);
	fmt.Println(fullNameNew);
	fmt.Println(age);
	fmt.Println(firstName + " " + lastName);
}