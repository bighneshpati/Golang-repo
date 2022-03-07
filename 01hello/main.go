package main

import "fmt"

func main() {
	fmt.Println("Hello to the go world")
	if (34 < 36) && (67 < 89) {
		fmt.Println("Hello to the new world")
	}
	for i := 0; i < 3; i++ {
		fmt.Println("hello world to the for loop")
	}
	i := 0
	for i <= 5 {
		if i == 3 {
			break
		}
		fmt.Println(i)
		i += 1
	}
}
