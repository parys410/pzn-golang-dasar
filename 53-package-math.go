package main

import (
	"fmt"
	"math"
)

/**
** Package yang berisi constant dan fungsi matematika
** https://golang.org/pkg/math/

** Yang mungkin sering digunakan:
** math.Round(float64)					=> Membulatkan float64 ke atas atau ke bawah, sesuai dengan yang paling dekat
** math.Floor(float64)					=> Membulatkan float64 ke bawah
** math.Ceil(float64)						=> Membulatkan float64 ke atas
** math.Max(float64, float64)		=> Mengembalikan nilai float64 paling besar
** math.Min(float64, float64)		=> Mengembalikan nilai float64 paling kecil
 */

func main() {
	float1 := 3.251284521
	float2 := 3.251432832

	fmt.Println(math.Round(float1))
	fmt.Println(math.Floor(float1))
	fmt.Println(math.Ceil(float1))
	fmt.Println(math.Max(float1, float2))
	fmt.Println(math.Min(float1, float2))
}