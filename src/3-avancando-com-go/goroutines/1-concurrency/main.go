package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	go runProcess("P1", 20)   // Rotina rodando
	go runProcess("P2", 20)

	var s string              // Variável que pega e armazena os valores do processo
	fmt.Scanln(&s)            // E o Scan verifica as alterações em s
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}

// Paralelismo
// P1 ___________________________________ (1 core) 
// P2 ___________________________________ (2 core)

// Concorrência(Concurrency)
// P1 ___      _________     ______    ____ (1 core) 
// P2 _________         _____      ____     (2 core)