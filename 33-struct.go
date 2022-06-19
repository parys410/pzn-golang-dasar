package main

import "fmt"

/**
** Struct
** Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
** Struct biasanya representasi data dalam program aplikasi yang kita buat
** Data struct disimpan dalam field
** Sederhananya struct adalah kumpulan dari field (Seperti Object di Javascript)
** Nama Struct biasanya PascalCase
** Agar bisa di export / digunakan di modul lain maka nama field menggunakan PascalCase

** How To Make Struct
type [NamaStruct] struct {
	Field1, Field2 string
	Field3 int
	Field4 bool
}

** How To Make Object Struct (Struct Literals)
*!Cara Pertama
var namaObject NamaStruct
namaObject.Field1 = "Value1"
namaObject.Field2 = "Value2"
namaObject.Field3 = 0
namaObject.Field4 = false

*!Cara Kedua
namaObject := NamaStruct {
	Field1: "Value1",
	Field2: "Value2",
	Field3: 0,
	Field4: false,
}

*!Cara Ketiga
*!Cenderung error apabila struktur/urutan field dari Struct diubah atau ada penambahan field
namaObject := NamaStruct{"Value1", "Value2", 0, false}

** Membuat Data Struct
** Struct adalah template data / prototype data / class di OOP
** Struct tidak bisa langsung digunakan
** Namun kita bisa membuat data/object dari struct yang telah kita buat
*/

type Customer struct {
	Name, Address string
	Age int
	Married bool
}

func main() {
	var pelanggan Customer
	pelanggan.Name = "Ary"
	pelanggan.Address = "Abianbase"
	// pelanggan.Age = 29

	pelanggan2 := Customer{
		Name: "Dwija",
		Address: "Denpasar",
		Age: 30,
		Married: true,
	}

	pelanggan3 := Customer{ "Bangkit", "Denpasar", 29, false }

	//* Jika terdapat field yang tidak didefinisi, maka secara default di reset ke nilai defaultnya
	//* Misalkan string = "", int = 0, bool = false

	fmt.Println("Pelanggan", pelanggan)
	fmt.Println("Pelanggan2", pelanggan2)
	fmt.Println("Pelanggan3", pelanggan3)
}