package main

import "fmt"

/**
** Defer
** Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
** Defer function akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi

** Panic
** Panic function adalah function yang bisa kita gunakan untuk menghentikan program
** Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
** Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

** Recover
** Recover adalah function yang bisa kita gunakan untuk menangkap data panic
** Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
** Pastikan Recover selalu disimpan di defer function
**/

func logging() {
	message := recover() //*! Biasakan recover itu selalu dieksekusi di akhir program, disini akan ditangkap error dari PANIC
	if message != nil {
		fmt.Println("Terjadi error", message)
	}
	
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() //*! Biasakan menggunakan defer itu di paling atas
	
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
	logging() //*! logging() tidak jalan karena sudah dihentikan oleh panic error
}

func runApp(error bool) {
	defer logging()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	// runApplication(0) //* for testing defer
	runApp(true) //* for testing defer, panic & recover
	fmt.Println("End main()") //* Akan tetap dieksekusi, karena sudah ada recover function
}