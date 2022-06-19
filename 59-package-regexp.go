package main

import (
	"fmt"
	"regexp"
)

/**
** Package regexp
** adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
** Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
** https://github.com/google/re2/wiki/Syntax
** https://golang.org/pkg/regexp/

** Yang mungkin digunakan
** regexp.MustCompile(string)			   => Membuat Regexp object
** Regexp.MatchString(string) bool   => Cek apakah Regexp match dengan string (dari Regexp Object)
** Regexp.FindAllString(string, max) => Cari string yang match dengan maximum jumlah hasil (dari Regexp Object) // max isi -1 untuk menemukan sebanyak-banyaknya
 */

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("EKO"))
	fmt.Println(regex.MatchString("ary"))

	search := regex.FindAllString("eko eka edo eto ary eyo eki", -1)
	fmt.Println(search)
}