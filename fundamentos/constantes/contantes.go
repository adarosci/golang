package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * m.Pow(raio, 2)
	fmt.Println("area bla", area)

	const (
		a = 1
		b = 2
	)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "a"

	fmt.Println(g, h, i)
}
