package main

import (
	"runtime"
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var waitGroup sync.WaitGroup

// Função que é executada logo que o sistema vai começar a rodar para fazer um setup da aplicação
func init() {
									       //Neste caso com 1 ele roda de forma concorrente
	//runtime.GOMAXPROCS(1)                // Seta a quantidade de cores que queremos que a aplicação use para ser executada
	
										  // Roda de forma paralela -> a partir da versão 1.5 isso não é necessário	
	runtime.GOMAXPROCS(runtime.NumCPU())  // Quando o Go for rodar, ele roda com o número máximo de CPU que temos

}

func main() {
	//fmt.Println(runtime.NumCPU())  Mostra a quantidade de CPUS que eu tenho na minha máquina para rodar os processos
	waitGroup.Add(2)              // Indica que tenho 2 processos para rodar 
	go runProcess("P1", 20)       // Processo 1
	go runProcess("P2", 20)       // Processo 2
	waitGroup.Wait()              // Aguarde os processos terminarem
	
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()            // Informe que o processo terminou
}