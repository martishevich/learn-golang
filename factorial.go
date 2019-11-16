package main

import "fmt"

func factorial(number uint) uint {
	result := uint(1)
	for i := uint(1); i <= number; i++ {
		result *= i
	}
	return result
}

func main() {
	result := factorial(5)
	fmt.Println(result)
}
