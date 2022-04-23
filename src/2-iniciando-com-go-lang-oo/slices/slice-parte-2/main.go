package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	//slice = append(slice, 10, 2, 5, 40)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// for i := 0; i < 20; i++ {
	// 	slice = append(slice , i)
	// 	fmt.Println("Len: ", len(slice), " Cap: ", cap(slice))
	// } 
	fmt.Println("------------------------------------------------")	

	sliceTeste := slice //  Isso não cria uma cópia e sim um apontamento para o mesmo endereço de memória
	slice[0] = 10 // Ao fazer isso estou alterando os dois slices, o slice e o sliceTest, pois ambos apontam para o mesmo local
	slice = append(slice, 1, 2, 3) 
	fmt.Println(slice) // -> imprime: [10 0 1 2 3] -> apesar de ter mais elementos do que o sliceTest o apontamento ainda é o mesmo
	fmt.Println(sliceTeste) // -> [10 0]

	fmt.Println("------------------------------------------------")	

	slice = append(slice, 1, 2, 3, 4, 5, 6) 
	fmt.Println(slice) 
	fmt.Println(sliceTeste) 
}