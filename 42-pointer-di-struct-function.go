package main

import "fmt"

/**
** Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
** Sangat direkomendasikan menggunakan pointer di method, sehingga tidak BOROS MEMORY karena harus selalu diduplikasi ketika memanggil method
 */

type Man struct {
	Name string
	IsMarried bool
}

func (man *Man) SetMarried() {
	man.IsMarried = true
	fmt.Println("In Function", man)
}

func main() {
	person1 := Man {
		Name: "Dede",
	}
	person1.SetMarried()
	fmt.Println(person1)
}