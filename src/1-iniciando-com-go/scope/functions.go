package scope

import "fmt"

var anoAtual int = 2022

func CalculaIdade(dataNascimento int) {
	fmt.Println("Você tem", (anoAtual - dataNascimento), "anos de idade")
}