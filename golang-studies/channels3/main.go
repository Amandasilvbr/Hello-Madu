package main

import (
	"fmt"
	"time"
)

func main() {
	hello := make(chan string) //criando channel

	go func ()  {
		time.Sleep(time.Second * 2) // dois segundos
		hello <- "hello world"
	}()

	select { // isso vai parar a execucao e ver se ta vindo algum valor pro x do hello
	case x:= <- hello:
		fmt.Println(x)
	default: fmt.Println("default")
	}

	time.Sleep(time.Second * 5)
}