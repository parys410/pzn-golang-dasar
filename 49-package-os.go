package main

import (
	"fmt"
	"os"
)

/**
** Package OS berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independedn (bisa digunakan di semua sistem operasi)
** https://golang.org/pkg/os/
** cara menggunakan argument adalah
** go run namafile.go args1 args2 args3 dst
 */

func main() {
	//! ARGUMENT
	args := os.Args
	fmt.Println("Arguments   :", args)

	//! HOSTNAME
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hostname    :", name)
	}

	//! GET ENVIRONMENT VARIABLES
	env := os.Getenv("DLC")
	fmt.Println("Environment :", env)
}