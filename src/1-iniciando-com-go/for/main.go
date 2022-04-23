package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// No Go não existe While, então temos:
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Gerando looping infinito
	for {
		x++

		fmt.Println(x)
		if x == 50 {
			continue
		}

		if x == 100 {
			break
		}
	}

}