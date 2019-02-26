package main

import (
	"fmt"
	"runtime"

	"./gorotine"
)

/*
Concorrencia e paralelismo
go e uma linguagem concorrente
paralelismo processar dois dados em processadores diferentes
concorrencia rodar varios processos no mesmo processador
*/
func main() {
	fmt.Println(runtime.NumCPU())

	ch := make(chan gorotine.Pessoa, 30)
	p := gorotine.Pessoa{"Alisson", 2}

	go p.Fale(50, ch)

	for i := range ch {
		fmt.Println(i)
	}
}
