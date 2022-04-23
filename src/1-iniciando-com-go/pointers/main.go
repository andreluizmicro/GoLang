package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	return *a
}

func main() {

	b := 10
	fmt.Println(xpto(&b))
	fmt.Println(b)

	x := 10

	// & printa o endereço de memória de x
	fmt.Println(&x)
	y := &x

	fmt.Println(y) // Printa o endereço de memória 0xc420098010
 
	fmt.Println(*y) // Printa o valor 10 que está dentro do endereço de memória 0xc420098010

	*y = 20
	fmt.Println("x de novo: ", x) // x passa a ser 20 também

	var z *int = &x
	fmt.Println(*z)

}
