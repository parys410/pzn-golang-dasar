package main

import "fmt";

func main(){
	var nama1 string = "Ary";
	var nama2 string = "Setiyawan";
	var number1 int8 = 10;
	var number2 int8 = 5;

	var checkNama bool = nama1 != nama2;
	var checkNumber bool = number1 >= number2;
	fmt.Println(checkNama, checkNumber);
}