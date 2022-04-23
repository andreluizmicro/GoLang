package main

import (
	"sync"
	"fmt"
)

func main() {

	channel := make(chan int)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func(){
		waitGroup.Wait()
		close(channel)          // Uma vez fechado não é mais possível mecher nesse channel
	}()

	for number := range channel {
		fmt.Println(number)
	}
	
	
}
