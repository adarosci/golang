package main

import "fmt"

type esportivo interface {
	logarTurbo()
}

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var a imprimivel = pessoa{"nome", "sobrenome"}
	fmt.Println(a.toString())
	a = produto{}
	imprimir(a)
	mainInterface()
}

//tipo interface
type curso struct {
	nome string
}

func mainInterface() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = "a"
	fmt.Println(coisa)

	coisa = curso{"Joao"}
	fmt.Println(coisa)
}
