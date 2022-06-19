package main

import "fmt"

/**
** Go-Lang akan memberikan nilai default pada saat buat variable baru, int=0; string=""; bool=false
** Go-Lang memiliki data nil, yaitu data kosong
** Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
 */

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string] string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Ary")
	person2 := NewMap("")

	if person2 == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person2)
	}
	
	fmt.Println(person)
}