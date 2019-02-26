package main

import "fmt"

type carro struct {
	nome       string
	valocidade int
}

type ferrari struct {
	carro
	turbo bool
}

type nota int

func (n nota) entre(i, f int) bool {
	return int(n) >= i && int(n) <= f
}

func main() {
	f := ferrari{}
	fmt.Println(f)
	var aa nota
	aa = 1
	fmt.Println(aa)
}
