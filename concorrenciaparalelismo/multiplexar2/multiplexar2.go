package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second)
			c <- "Oi eu sou " + pessoa
		}
	}()
	return c
}

func juntar(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("joao"), falar("maria"))
	for i := range c {
		fmt.Println(i)
	}
}
