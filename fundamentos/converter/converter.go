package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	//converter int para string
	fmt.Println("teste" + strconv.Itoa(123))
	fmt.Println("teste", 123)
	fmt.Println("teste" + string(97))

	num, _ := strconv.Atoi("123")

	fmt.Println(num)
}
