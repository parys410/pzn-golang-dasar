package main

import (
	"fmt"
	"time"
)

/**
** Package Time
** adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
** Tanggal, Bulan, Tahun, Jam, Menit, dan Detik
** https://golang.org/pkg/time/

** Yang Mungkin biasa digunakan:
** time.Now()									=> Untuk get waktu saat ini
** time.Date(...)							=> Untuk membuat waktu
** time.Parse(layout, string)	=> Untuk memparsing waktu dari string
 */

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	utc := time.Date(2022, 6, 19, 19, 43, 0, 0, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, err := time.Parse(layout, "2020-10-25")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		newDate := parse.Format("02/01/06")
		newDate2 := parse.Format("02/Jan/2006")
		fmt.Println(parse, newDate, newDate2)
	}
}