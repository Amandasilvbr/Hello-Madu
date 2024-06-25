package main

import "fmt"

// um channel serve p fazer uma thread se comunicar com a outra

//thread 1
func main() {

	// isso aqui é a comuniccao da thread 1 com a 2
	hello := make(chan string) // criando um canal e atribuindo a um variavel

	//thred 2
	go func ()  { // isso é uma go routine
		hello <- "Hello world"
	}()

	result := <-hello // toda vez q tiver um valor em hello, envia pra result
	fmt.Println(result)
}