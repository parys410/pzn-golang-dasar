package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	censoredName := ""
	if name == "Anjing" {
		for i, word := range name {
			if i > len(name) / 2 {
				censoredName += "*"
			} else {
				censoredName += string(word)
			}
		}
	} else {
		censoredName = name
	}
	return censoredName
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Ary Setiyawan", spamFilter)
}