package main

import "fmt"

//Aula - Slices - São arrays um pouco mais dinânimicos

func main()  {

	// make -> produz ou cria um slice na memória
	slice := make([]int, 5)
	slice = append(slice, 1, 2, 3)

	fmt.Println(slice)
	fmt.Println(len(slice))

	fmt.Println("----------------------------")
	slice2 := make([]int, 2, 5) // Tem duas posições e capacidade 5

	// Causa erro!!!
	// slice2[2] = 10
	// slice2[3] = 20

	slice2 = append(slice2, 10, 2, 6, 40, 50) // Ao passar da capacidade 5 ele dobra para 10, com isso a capacidade passa a ser 10

	fmt.Println(slice2)
	fmt.Println("Tamanho:", len(slice2))
	fmt.Println("Capacidade:", cap(slice2))


}
