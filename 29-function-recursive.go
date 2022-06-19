package main

import "fmt"

func loopFactorial(angka int) int {
	result := 1
	for i := angka; i >= 1; i-- {
		result *= i
	}
	return result
}

func recursiveFactorial(angka int) int {
	if angka > 1 {
		return angka * recursiveFactorial(angka - 1)
	} else {
		return angka
	}
}

func main() {
	angka := 8
	fmt.Printf("Factorial from %d is %d \n", angka, loopFactorial(angka))
	fmt.Printf("Factorial from %d is %d", angka, recursiveFactorial(angka))
}