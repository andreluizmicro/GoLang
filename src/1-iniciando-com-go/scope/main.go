package scope

import "fmt"

var y int = 20

// func main() {
// 	x := 10
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	printX(50)
// 	fmt.Println(author)
// 	printZ()
// }

func printX(num int) {
	fmt.Println("y +", num , " é igual a: ", y+num)
}

var z int = 1000