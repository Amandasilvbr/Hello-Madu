package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome string
	Email string
	CPF int
}

func (c Cliente) Exibe() { // criei esse metodo e vou associar a uma struct (como se fosse a relacao classe - metodos)
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)
}

type ClienteInternacional struct { // essas duas structs sao praticamente iguais, portanto eu poderia declarar cliente la dentro e declrar so oq tiver de diferente
	Cliente // composicao de struct
	Pais string `json:"pais"` // transformando pra letra minuscula
}

func main() {
	
	cliente := Cliente {
		Nome: "madu",
		Email: "madu@gmail.com",
		CPF: 1234567,
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"maria", "m@gmail.com", 123}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional {
			Cliente: Cliente{
			Nome: "davi",
			Email: "d@d.com",
			CPF: 1234,
		},
		Pais: "Alemanha", 
	}
	fmt.Printf("Nome: %s. Email: %s. Pais: %s\n ", cliente3.Nome, cliente3.Email, cliente3.Pais) // mesmo fazendo a composicao de uma struct com outra eu consigo fazer 

	// observando heranca: os metodos de uma struct sendo aplicados sobre a struct que utiliza ela 
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	// pegando uma struct e convertendo ela pra json
	clienteJson, err := json.Marshal(cliente3) 
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clienteJson)) // por natureza isso retorna um slice de bytes, por isso convetemos p string


	// populando uma stuct com json
	jsonCliente4 := `{"Nome":"davi","Email":"d@d.com","CPF":1234,"pais":"Alemanha"}`
	cliente4 := ClienteInternacional{}
	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}