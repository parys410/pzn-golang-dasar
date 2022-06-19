package main

import "fmt"

type any interface{}

func Ups(i int) any {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data any = Ups(2)
	data2 := Ups(1)
	fmt.Println(data, data2)
}