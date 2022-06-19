package main

import "fmt";

func main(){
	/*
		+ : Penjumlahan
		- : Pengurangan
		* : Perkalian
		/ : Pembagian
		% : Sisa Pembagian (Mod)

		AUGMENTED ASSIGNMENT
		a = a + 10	: a += 10
		a = a - 10	: a -= 10
		a = a * 10	: a *= 10
		a = a / 10	: a /= 10
		a = a % 10	: a %= 10

		UNARY OPERATOR
		a = a + 1		: a++
		a = a - 1		: a--
		negative		: -
		positive		: + (tidak wajib)
		kebalikan		: ! (not for boolean)
	*/

	var a = 10;
	var b = 10;
	var c = a + b;
	var d = a / 3;
	var e = a % 3;
	var d2 float32 = float32(a) / 3;
	a += 100;
	a++;
	var f bool = true;
	var g = !f;

	fmt.Println(a, b, c, d, e, d2, f, g);
}