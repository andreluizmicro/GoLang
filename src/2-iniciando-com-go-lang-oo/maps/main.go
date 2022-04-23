package main

import "fmt"

// Os maps no Go podem ser comparados a arrays associativos no PHP ou dicionário no Python...
// Os maps funcionam como variáveis que tem key e value

func main() {

	m := make(map[string]int) // O indice é uma string e o valor um inteiro
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m, "c")
	fmt.Println(m["c"])

	//
	_, exists := m["c"] 
	fmt.Println(exists)

	value, exists := m["a"]
	fmt.Println(value)

	// var x = map[string]int{};
	// fmt.Println(x)

	newMap := map[string]int{"x":5, "y":10}
	fmt.Println(newMap)

	if value, exists := m["c"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("ops!")
	}

}