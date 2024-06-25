package main

import "fmt"

type Number interface { // criando meu proprio tipo
	~int | int64 | float64 | float32
}

type MyNumber int // se tiver o ~ ali em cima, significa q myNumber implementa a interface Number, o que significa que os numeros desse tipo sao aceitos

func SomaInteiros(m map[string]int64) int64{
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

/* isso é comum fazer, mas erro. repetir tantas linhas de codigo apenas pq o tipo é diferente.
func SomaFloat(m map[string]float64) float64{
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}*/

// eu poderia adaptar para fazer da melhor: assim, usando generics
func SomaGenerica[T int64 | float64] (m map[string]T) T{
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

//ou usando o tipo que criei, asssim (melhor forma)
func SomaGenerica2[T Number] (m map[string]T) T{
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// usando o conceito de generics em uma operacao de soma comum.
func Soma[T Number](number1 T, number2 T) T {
	return number1 + number2
}
// um problema que eu podeia ter ao usar generics criando um tipo seria na comparacao. se eventualmente eu tentasse comparar algo que vier para mim 
// (pode ser de qualquerr tipo, eu nao sei qual) com um tipo definido, teria problema pois um tipo criado é incomparavel. 
func Soma2[T comparable](number1 T, number2 T) T { // usando o comparable eu consigo comparar igualdade. agora se fosse number1 > number2 eu ja nao conseguiria. 
	//existe um pacote go chamado constraints que nos ajudaria a fazer isso
	if number1 == number2 {
		return number1
	}
	return number2
}

func main(){
	//var x, y, z MyNumber 
	//x = 1
	//y = 2
	//z = 3


	fmt.Println(SomaGenerica2(map[string]MyNumber{"x": 1, "y": 2, "z":3}))
	//fmt.Println(SomaFloat(map[string]float64{"a": 1.1, "b": 1.2, "c":2.3}))
}