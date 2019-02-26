package main

import "fmt"

func valor(n int) {
	defer fmt.Println("FIm")
	fmt.Println(n)
}

func main() {
	valor(4)
}
