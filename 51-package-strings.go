package main

import (
	"fmt"
	"strings"
)

/**
** Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data string
** https://golang.org/pkg/strings/
** Yang mungkin sering digunakan:
** strings.Trim(string, cutset)					=> Memotong cutset di awal dan akhir string
** strings.ToLower(string)							=> Ubah semua karakter menjadi lowercase
** strings.toUpper(string)							=> Ubah semua karakter menjadi UPPERCASE
** strings.Split(string, operator)			=> Memotong string berdasarkan separator
** strings.Contains(string, search)			=> Cek apakah string mengandung string lain
** strings.ReplaceAll(string, from, to)	=> Ubah semua string dari from ke to
 */

func main() {
	kalimat := " Putu Ary Setiyawan "
	trim := strings.Trim(kalimat, " ")

	fmt.Println(kalimat, len(kalimat))
	fmt.Println(trim, len(trim))
	fmt.Println(strings.Contains(kalimat, "ary"))
	fmt.Println(strings.Contains(strings.ToLower(kalimat), "ary"))
	fmt.Println(strings.Split(trim, " "))
	fmt.Println(strings.ReplaceAll(trim, "Ary", "Eko"))
}