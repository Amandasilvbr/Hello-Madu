package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() { // essa func é uma thread

	runtime.GOMAXPROCS(1) // nosso processador é multicore, ele ode acabar consgeuindo rodar em mais de um. entao definimos para rodar somente em um core
	fmt.Println("Comecou")

	/* o Go por padrao trabalha de forma coopertiva, mas a partir da versao 1.14, ele adicionou alguns recursos no schedule que quando aocntecem situcoes como
	essa (uma thread demorar exarcebadamente para terminar) ele simplesmente interrompe ela e coloca outra no lugar, ou seja, consegue trabalhar tbm de forma
	preemptiva

	go func() { // essa tbm é thread
		for { //loop infinito

		}
	} ()
	*/

	time.Sleep(time.Second)
	fmt.Println("Terminou")
}