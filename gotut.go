package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

// func subtract(a, b uint64) uint64 {
// 	return a - b
// }
func main() {
	var (
		id          int    = 23
		name        string = "Bighnesh"
		nameAnother string = "Achintya"
		boolean     bool   = true
	)
	fmt.Println("The values are", id, name, nameAnother, boolean)
	num1, num2 := 4.6, 5.6
	fmt.Println("Hello world how are you ?")
	fmt.Println(add(num1, num2))
	fmt.Printf("%T, %v", num1, num1)
	// fmt.Println(subtract(num1, num2))
}
