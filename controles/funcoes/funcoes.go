package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2(p string, b string) {
	fmt.Println(p, b)
}

func f3() string {
	return "f3"
}

func f4(a, b, c string) string {
	return "f4"
}

func f5() (string, string) {
	return "r1", "aa"
}

func f6(a, b int) (ba int, aa int) {
	ba = b
	aa = a
	return
}

var soma = func(a, b int) int {
	return a + b
}

func multiplacacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, a, b int) int {
	return funcao(a, b)
}

func multiplosParametro(numeros ...int) int {
	return numeros[0]
}

func funcaoMagrinha(aprovados ...string) {

}

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

//recursivo
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("erro")
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

//ponteiros
func inc2(n *int) {
	*n++
}

func main() {
	x := 1
	inc2(&x)

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(1, 2))

	result := exec(multiplacacao, 2, 3)
	fmt.Println(result)

	ap := []string{"a", "b", "c"}
	funcaoMagrinha(ap...)

	//ap1 := [3]string{"b", "c", "f"}

	//nao funciona com array
	//funcaoMagrinha(ap1)

	r, _ := fatorial(5)
	fmt.Println(r)
}
