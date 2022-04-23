package main

import ( 
	"fmt"
)


func main() {

								// msg é um channel 
	msg := make(chan string)	// make -> função especifica para criar variáveis dos tipos: slice, channels, map

	// Função anônima
	go func() {
							 	// Forma de consegui jogar um valor para o channel mensagem
		msg <- "Hello world" 	// Mande para o meu channel uma mensagem Hello world
	}()
								// Precisamos liberar o channel para que possa receber mais mensagens, para isso esvaziamos ela 
	result := <-msg				// Fazendo o channel passar a mensagem para frente(ou para alguém) - Esvaziamos o channel msg
	fmt.Println(result)

}