package main

import "fmt"

/**
** Pointer di Function
** Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
** Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah
** Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
** Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
** Untuk mealkukan ini, kita juga bisa menggunakan pointer di function
** Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
 */

type DatabaseConfig struct {
	DatabaseName, Username, Password string
	Port int
}

func ResetDatabase(db *DatabaseConfig) {
	db.DatabaseName = "localdb"
	db.Username = "root"
	db.Password = "root"
	db.Port = 3036

	fmt.Println("Reset Database", db)
}

func main() {
	db1 := DatabaseConfig{
		DatabaseName: "belajar",
		Username: "admin",
		Password: "admin",
		Port: 8080,
	}

	ResetDatabase(&db1)

	fmt.Println(db1)
}