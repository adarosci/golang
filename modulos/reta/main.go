package main

import "./area"
import "fmt"

func main() {
	p1 := area.Ponto{1.5, 2.3}
	p2 := area.Ponto{2.5, 3.3}
	fmt.Println(area.Distancia(p1, p2))
}
