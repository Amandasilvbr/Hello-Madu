package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i:= 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {

	/*
	contador("sem go routine")
	go contador("com go routine") // go na frente cria um thread, essa atua de forma concorrente
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
	time.Sleep(time.Second)
	fmt.Println("fim...")
	*/

	go contador("a")
	go contador("b")
	time.Sleep(time.Second * 10) // executando de forma concorrente
}