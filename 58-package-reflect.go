package main

import (
	"fmt"
	"reflect"
)

/**
** Package Reflect (reflection)
** Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
** Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
** Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
** https://golang.org/pkg/reflect/
 */

type Sample struct {
	Title string `required:"true" max:"10"` //! StructTag
}

type ContohLagi struct {
	Name string `required:"true"` //! StructTag
	Description string `required:"true"` //! StructTag
}

func IsValid(data interface{}) bool {
	dataType := reflect.TypeOf(data)
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{ "Ary" }
	contoh := ContohLagi{ "Ary", "" }
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min"))

	sample.Title = ""

	fmt.Println(IsValid(contoh))
}