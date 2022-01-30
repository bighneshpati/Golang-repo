package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza")
	// comma ok syntax // err ok
	input, error := reader.ReadString('\n')
	fmt.Println("Thanks for reading, ", input)
	fmt.Printf("Type of the rating is %T", input)
	fmt.Println()
	fmt.Println(error)
}
