package main

import "fmt"

func main() {

	forever := make(chan string) //criando um canal

	go func() {
		
	}()
	fmt.Println("aguardando pra sempre")
	<-forever // enquanto o forever nao receber nenhum valor, ele vai travar a minha execucao
}