package main

import "fmt"

func funcMultiplica(a int) int {
	return a * a
}

// Cuidado ao usar isso por conta de legibilidade
func namedReturn(a string) (x string) {
	x = a
	return
}

// Em Go podemos ter uma função com mai de um retorno
func moreReturns(a string, b string) (string, string) {
	return a, b
}

// Função variádica
func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}


// Função que retorna uma função que retorna um inteiro
func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {

	x := funcMultiplica(5)

	fmt.Println(x)
	fmt.Println(funcMultiplica(10))

	fmt.Println("------------------------")
	fmt.Println(namedReturn("André"))

	fmt.Println("------------------------")
	nome, sobrenome := moreReturns("André", "Luiz")

	fmt.Println(nome, sobrenome)

	fmt.Println("------------------------------------------")
	fmt.Println("---------- FUNÇÃO VARIÁDICA --------------")
	fmt.Println("------------------------------------------")
	fmt.Println(variadicFunc(1, 2, 5, 10))

	// Função Anônima
	z := 0

	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
	fmt.Println(add())
}