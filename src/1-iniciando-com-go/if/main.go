package main

import "fmt"

func main() {

	a := 4

	if a > 10 {
		fmt.Println("A > 10")
	} else if (a > 5) {
		fmt.Println("A > 5 e menor que 10")
	} else {
		fmt.Println("Nenhuma condição foi aceita.")
	}
	
	b := true

	if b {
		fmt.Println("Sim B é true!")
	}

	// em tempo real x é igual a cool, se b for true imprima
	if x := "Cool"; b {
		fmt.Println(x)
	}
}