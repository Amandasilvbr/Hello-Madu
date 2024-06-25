package main

import (
	//"fmt"
	//"curso-go/math"  importando pacote criado por mim
	//"errors"
	//"fmt"
	//"log"
	// "net/http"
)


/* POO EM GO?

type Carro struct { // os metodos em go sao basedos nas structs, assim se trabalha com POO em go, ou "Go Way" 
	Nome string
}

func(c Carro) andar() { // desta forma eu relaciono esta minha funcao ao struct
	fmt.Println(c.Nome, "andou")
}
*/

func main() { 

	/*
	carro := Carro{
		Nome: "BMW",
	}

	carro.andar()
	*/
	

/* VARIAVEIS, TIPOS E PACOTES

	a := "madu" // declaracao e atribuicao, go infere o tipo
	// este tipo nao pode ser mudado neste mesmo formato

	// TIPOS
	b:= 10
	c:= 3.144
	d:= false

	fmt.Printf("%v \n", a) // v retorna valor
	fmt.Printf("%v \n", b) 
	fmt.Printf("%T \n", c) // T retorna type
	fmt.Printf("%T \n", d)


	resultado := math.Soma(1, 1) // go inferindo o tipo
	fmt.Printf("%T", resultado) 
*/


/* TRATANDO ERROS

	res, err := http.Get("http://godggthrtjhrnyjoglecom.br") // isso retorna duas variveis:  resposta e error, se nao houver erro esta ultima vem vazia. grade prt das funcoes em go segue esse padrao
	if err != nil {
		log.Fatal(err.Error()) // log de erro
	}

	fmt.Println(res.Header) // imprimindo o header da requisicao

	
	//testando a soma1 
	res, err := soma1(7, 2) 
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
*/


/* FUNCOES
	
	retornando dois valores na funcao soma2
	resultado, str := soma3(10, 20)
	fmt.Println(resultado, str)

	testando a funcao somaTudo
	resultado := somaTudo(3, 5, 8, 9, 15)
	fmt.Println(resultado)



	resultado := func(x ...int) func() int { 
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}
*/


/* PONTEIROS

	// memoria-XYZ(50) <---- a <----- 50 // a variavel aponta p um lugar na memoria, que tem um endereco. quando eu atribuo um valor a variavel, o mesmo é armanzenado neste endereco
	a := 10
	fmt.Println(&a) // isso me retorna onde essa variavel esta sendo guardada

	var ponteiro *int = &a // essa variavel ponteiro sabe onde o "a" está armazenando o valor. ela detem o endereco de memoria
	fmt.Println(*ponteiro) // retorna o valor armazenado no endereco de memoria. sem o aterisco retorna o endereco de memoria

	*ponteiro = 50 
	fmt.Println(*ponteiro) // consigo mudar o valor armazenado no endereco de memoria pelo ponteiro


	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)// isso vai retornar 200, pois a func abc altera o valor
	*/

// STRUCTS, COMPOSICAO E JSON

}


/* TRATANDO ERROS
	func soma1 (x, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil 
}
*/

/* FUNCOES

	func soma2(a int, b int) { // isso nao funciona pq o go exige uma execificacao do tipo de retorno
	return a + b
}

func soma3(c int, d int) (int, string){ // no go eu posso retonar mais do que um valor
	return c + d, "somou"
}

func somaTudo(x ...int) int {
	resultado := 0 

	for _, v := range x { // fazendo um loop de todos os valores em x, no caso o x é como um array que armzena varios valores
		resultado += v
	}

	return resultado
}
*/

/* PONTEIROS
func abc(a *int) int {
	*a = 200
}
*/


