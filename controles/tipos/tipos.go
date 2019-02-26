package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

//metodo
func (p pessoa) getIdade() int {
	return p.idade
}

func (p *pessoa) setName(n string) {
	p.nome = n
}

func main() {
	p := pessoa{
		"Joao", 10,
	}
	fmt.Println(p, p.getIdade())

	p.setName("Alisson")
	fmt.Println(p)

}
