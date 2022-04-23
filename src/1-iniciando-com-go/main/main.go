package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	u := uuid.New().String()
	fmt.Println("Gerado UUID: ", u)
}