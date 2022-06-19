package main

import (
	"flag"
	"fmt"
)

/**
** Kadang-kadang kita ingin mendapat argument dengan benar
** Package flag berisikan fungsionalitas untuk memparsing command line argument
** Untuk menjalankannya:
** go run namafile.go -argt1=value1 -argt2=value2 -argt3=value3 dst...
** Balikannya selalu pointer *
 */

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 100, "Put your number")
	flag.Parse()

	fmt.Println(*host, *username, *password, *number)
}