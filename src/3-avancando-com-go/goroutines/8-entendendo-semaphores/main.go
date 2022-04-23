package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	ok := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i	
		}
		ok <- true               // Quando ok recebe true, ele está cheio e terá que ser esvaziado
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i	
		}
		ok <- true
	}()

	go func() {
		<-ok // Esvazia ele para poder fechar com o close - ou seja flag que confirma que foi esvaziado
		<-ok 
		close(channel) // Fecha o canal somente 
	}()

	for number := range channel {
		fmt.Println(number)
	}
}