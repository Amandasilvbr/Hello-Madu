package math

var A string = "showww" // consigo ter acesso a essa variavel no pckge main utilizando math.A

func Soma(a, b int) int { // funcao tem q comecar com letra maiuscula para ser exportada
	return a + b
}