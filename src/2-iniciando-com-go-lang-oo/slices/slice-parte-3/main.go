package main

import "fmt"

func main() {

	//sliceString := [10]string { Decalarando dessa forma passa a ser um array e não um slice

	sliceString := []string {
		"Hello",
		"World",
		"Much",
		"Better",
		"Better 2",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0]) // Hello
	fmt.Println(sliceString[:2]) // Hello World pega as posições de 0 a 2
	fmt.Println(sliceString[1:2]) // World
	fmt.Println(sliceString[2:4]) // Much Better
	fmt.Println(sliceString[2:]) // Pega do dois ao infinito

}