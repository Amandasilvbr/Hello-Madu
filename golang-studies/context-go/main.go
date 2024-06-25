package main

import (
	"context"
	"fmt"
	"time"
)

/* se um microsservico demora muito p responder, por ex 2 seg, context consegue interromper
esse pacote no fim das contas, dita as regras de tempo limite que as aplicacoes vao suportar ficar sem resposta
o context acaba sendo um carregador de informacoes tbm, e consigo utiliza-lo p interromper processamento na hora
*/

// contexto funciona como arvore, posso ter um ai e varios filhos, um dentro do outro. assim se eu emito sinal em um contexto, consigo
// ir parando outros contextos

//vou fazer um sistema de reserva de hotel
func main() {
	ctx := context.Background() //contexto pai, raiz, em branco
	ctx, cancel := context.WithCancel(ctx) // gerando um novo contexto q retorna uma funcao de cancelamento. quando eu executar essa funcao de canc todo mundo q tiver esse contexto vai receber
	// um sinal avisando q foi abortado
	defer cancel() // defer espera tudo ser executado e executa por ultimo

	go func ()  { //crtiando outra thread, ou seja isso vai ficar rodando em backgroung
		time.Sleep(time.Second * 4)
		cancel()
	}()

	bookHotel(ctx)
}

// aqui se o sistema demorar muito vai abortar e se nao, eu consigo fzer minha reserva
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // ja foi cancelado
		fmt.Println("Temo excedido para bookar o quarto")

	case <-time.After(time.Second * 5): // nao cancelou em 5 segundos
		fmt.Println("Quarto reservado com secesso")
	}
}