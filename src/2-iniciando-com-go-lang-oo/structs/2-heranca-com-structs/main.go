package main

import "fmt"

// Structs - Forma como o GO trabalha com a orientação a objetos

type Car struct {
	Name string
	Year int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corolla", 2017, "White"}
	fmt.Println(car1.info())
	
	superCar := SuperCar{
		Car: Car{
			"Fusca",
			2030,
			"Blue",
		},
		CanFly: true,
		Name: "Elantra", // Substitui o valor inicial
	}

	fmt.Println(superCar)
	fmt.Println(superCar.Name)
	fmt.Println(superCar.info())

}