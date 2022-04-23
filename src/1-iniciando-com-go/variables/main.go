package main

import "fmt"

var b int	
var nome = "André Luiz"
var c, d string = "Letra c", "letra d"

const xvz int = 1333

const (
	Visible string = "Ola" // Letra maiscula deixa visivel fora deste pacote
	aa string = "Ola mundo"
	bb = 66
	cc = 77
)

func main() {

	// Constantes

	const xpto = 10


	// Variáveis
	a := 20
	product := "Smartphone"
	description := "Excelente telefone para uso pessoal"
	preco := 2.500
	touchscreen := true
	memoria := "4GB"
	nucleos := 'N'
	backtick := `Opa e ai tudo bem
				como vai?			
	`

	fmt.Printf("%v \n", a)
	fmt.Println("------ Informações do produto --------")
	fmt.Printf("%v \n", product)
	fmt.Printf("%v \n", description)
	fmt.Printf("%v \n", preco)
	fmt.Printf("%v \n", touchscreen)
	fmt.Printf("%v \n", memoria)
	fmt.Printf("%v \n", nucleos)
	fmt.Printf("%v \n", backtick)

	fmt.Println("------ Tipos das variáveis --------")
	fmt.Printf("%T \n", product)
	fmt.Printf("%T \n", description)
	fmt.Printf("%T \n", preco)
	fmt.Printf("%T \n", touchscreen)
	fmt.Printf("%T \n", memoria)
	fmt.Printf("%T \n", nucleos)

}
