package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[13] = "aaa"
	aprovados[222] = "bb"
	aprovados[3441] = "cc"
	fmt.Println(aprovados)

	for index, value := range aprovados {
		fmt.Println(index, value)
	}

	delete(aprovados, 13)
	fmt.Println(aprovados)

	map2 := map[string]float64{
		"a": 1.5,
		"b": 2.5,
	}
	map2["b"] = 3.4

	mapSubMap := map[string]map[string]int{
		"A": {
			"Alisson": 1,
			"Amanda":  2,
		},
		"B": {
			"Bruna":  3,
			"Bianca": 4,
		},
	}
	fmt.Println(mapSubMap)
}
