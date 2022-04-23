package main

import ( 
	"fmt"
	"time"
)

// Lembrar do Exemplo de Natação
func main() {

								
	channel := make(chan int)	

	// Função anônima roda em paralelo com a função 2
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i							// Aqui o looping fica parado até que channel seja esvaziado
		}						 	
	}()
		
	// Função anônima roda me paralelo com a função 1
	go func() {
		for {
			fmt.Println(<-channel)					// Esvazia o channel
		}
	}()
	time.Sleep(time.Second)

}