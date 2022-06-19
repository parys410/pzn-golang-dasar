package main

/**
** Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita di akses
** Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
** Untuk membuat function yang diakses secara otomatis ketika package di akses, kita cukup membuat function dengan nama init

** Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
** Secara default, Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan
** Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import
 */

import (
	_ "github.com/parys410/GoLang-Dasar/database"
)

func main() {

}