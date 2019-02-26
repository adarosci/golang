package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randon() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := randon(); i > 5 {
		fmt.Println(i)
	}
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	ai := 0
	for ai < 10 {
		ai++
	}

	switch ai {
	case 10:
		fmt.Println(ai)
	}

	switch {
	case ai == 3:
		fmt.Println('a')
	}

	fmt.Println(tipo('a'))
	fmt.Println(tipo(1))

	var notas [2]int
	notas[0], notas[1] = 7, 9
	fmt.Println(notas)

	numeros := [...]int{1, 2, 3, 4, 5}

	//indice, valor
	for i, n := range numeros {
		fmt.Println(i, n)
	}

	a1 := []int{3, 4, 5}
	fmt.Println(a1)

	//slice, pedaço do array
	s1 := a1[1:2]
	fmt.Println(s1)

	s2 := a1[:2]
	fmt.Println(s2)

	//slice make
	s := make([]int, 10)
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s, len(s), cap(s))

	//copy elements slice
	s3 := make([]int, 10)
	copy(s3, s)
	fmt.Println(s, s3)
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	default:
		return "nint"
	}
}
