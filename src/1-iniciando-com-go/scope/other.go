package scope

import "fmt"

var author string = "André Luiz"
var nome string = "Paulo Silva"

func PrintZ() {
	fmt.Println("Estou dentro do pacote other")
}

func CalculaPreco(preco int) {
	juros := 10
	valorTotal := preco * juros
	fmt.Println(valorTotal)
}