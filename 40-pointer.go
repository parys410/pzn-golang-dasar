package main

import "fmt"

/**
** POINTER **
** Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
** Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya
** Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
** Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

** Function New
** Go-Lang juga punya function new yang bisa digunakan untuk membuat pointer
** Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
 */

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{ "Subang", "Jawa Barat", "Indonesia" }
	address2 := address1
	address3 := &address1 //! address3 di pointer ke alamat address1
	address4 := &address1 //! address4 di pointer ke alamat address1 
	address5 := new(Address)
	address5.City = "Jakarta"

	address2.City = "Bandung"
	address3.City = "Malang" //! Pada saat ini address1.City dan address4.city pun akan berubah ke "Malang" karena address1 dan address3 mengarah ke alamat yang sama
	address3 = &Address{ "Denpasar", "Bali", "Indonesia" } //! Saat ini address3 diubah alamatnya ke alamat baru
	*address4 = Address{ "Makassar", "Sulawesi Tenggara", "Indonesia" } //! Saat ini address1 dan address4 yang mengarah pada alamat yang sama awalnya, dipindah ke alamat baru
	

	fmt.Println("Original", address1) //* address1 tidak berubah, karena value pada address1 di duplikasi ke address2 bukan di referensikan
	fmt.Println("Pass By Value",address2)
	fmt.Println("Pass By Reference and Change Memory Address", address3)
	fmt.Println("Pass By Reference and Multiple Change Memory Address", address4)
	fmt.Println("Address5", address5)
}