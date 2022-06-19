package database

import "fmt"

var connection string

func init() { //! Seperti Constructor
	connection = "MySQL"
	fmt.Println("init() Executed")
}

func GetDatabase() string {
	return connection
}