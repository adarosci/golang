package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(3200))

	//somente positivo
	//uint8 uint16 uint32 uint68
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	var c rune = 'a'
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)

	a := 'a'
	a++
	fmt.Println(string(a))
}
