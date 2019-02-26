package main

import (
	"fmt"
	"time"
)

func valoresCanal() <-chan int {
	c := make(chan int)
	go func() {
		index := 1
		for {
			time.Sleep(time.Millisecond * 100)
			c <- index
			index++
		}
	}()
	return c
}

func main() {
	c := valoresCanal()
	go func() {
		for vl := range c {
			fmt.Println(vl)
		}
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("Fim -------------------------")
	go func() {
		for vl := range c {
			fmt.Println(vl)
		}
	}()
	time.Sleep(time.Second * 10)

	fmt.Println("Fim >>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}
