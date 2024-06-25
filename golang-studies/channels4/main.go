package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int) // especie de fila com um canal de inteiros

	go func ()  {
		i:= 0
		for {
			time.Sleep(time.Second)
			queue <- i //so quando o queue for esvaziado que podera voltar pro loop, ele vai ficar esperando ser lido
			i++
		}
	}()

	//so esta imprimindo pq estou lendo
	for x:= range queue {
		fmt.Println(x)
	}
}