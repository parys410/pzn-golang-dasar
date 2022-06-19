package main

import (
	"container/list"
	"fmt"
)

/**
** Package container/list
** Adalah implementasi struktur data double linked list di Go-Lang
** https://golang.org/pkg/container/list/

** Linked List itu seperti rantai/ikatan di struktur data
** Mirip seperti array/slice tapi tidak ada index
** Double linked list artinya ada hubungan next dan previous
** Data yang paling ujung pasti nil
** Data nya dinamis banget, jadi jangan khawatir memasukkan data apapun. Namun tidak enaknya tidak punya index jadi harus di iterasi satu persatu
 */

func main() {
	data := list.New()
	data.PushBack("Ary") //! PushBack input data paling akhir | PushFront input data ke paling awal
	data.PushBack("Setiyawan")
	data.PushFront("Putu")
	data.PushFront(1)
	data.PushBack(true)

	for e := data.Front(); e != nil; e = e.Next() {
		if e.Value == "Ary" {
			fmt.Println("Removed")
			data.Remove(e)
		} else {
			fmt.Println(e.Value)
		}
	}
	// fmt.Println(data)
	// fmt.Println(data.Front().Value)
	//! data.Front() akan mengeluarkan alamat paling pertama dari linkedList nya
	//! .Value digunakan untuk mengambil value yang terdapat pada alamat linkedList tersebut
}