package main

import "fmt";

func main(){
	type NoKTP string;
	type Married bool;

	var ktpAry NoKTP = "5103020410920009";
	var isMarried Married = true;
	fmt.Println(ktpAry, isMarried);
}