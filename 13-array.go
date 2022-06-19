package main

import "fmt";

func main(){
	var names [3]string;
	var data [3]string;
	names[0] = "Putu";
	names[1] = "Ary";
	names[2] = "Setiyawan";

	var values = [3]int{1,2,3};
	
	fmt.Println(names, values);

	// Array Function
	/* 
		len(array) : untuk mendapatkan panjang Array
		array[index] : untuk mendapatkan data diposisi index
		array[index] = value : untuk set value pada array tsb
	*/

	fmt.Println(data, len(data));
}