package main

import (
	"fmt"
	"strconv"
)

/**
** String Conversion
** Untuk melakukan conversi dari string ke tipe data lain / dari tipe data lain ke string
** https://golang.org/pkg/strconv/
** Yang mungkin sering digunakan:
** strconv.parseBool(string)					=> Ubah string ke bool
** strconv.parseFloat(string)					=> Ubah string ke float
** strconv.parseInt(string)						=> ubah string ke int64
** strconv.FormatBool(bool)						=> ubah bool ke string
** strconv.FormatFloat(float, ...)		=> ubah float64 ke string
** strconv.FormatInt(int, ...)				=> ubah int64 ke string
 */

func main() {
	boolean, err := strconv.ParseBool("12345")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(boolean)
	}

	number, err := strconv.ParseInt("1000000000000", 10, 8) //! Hati Hati dengan Base number
	number2, _ := strconv.Atoi("1000000000000")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}
	fmt.Println(number2)

	numberStr := strconv.FormatInt(1000000000000, 8) //! Hati Hati dengan Base number
	numberStr2 := strconv.Itoa(1000000000000)
	fmt.Println(numberStr)
	fmt.Println(numberStr2)
}