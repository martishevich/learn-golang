package main

import "fmt"

func main() {
	factorialArgument := 5
	result := 1
	for i := 1; i <= factorialArgument; i++ {
		result *= i
	}
	fmt.Println(result)
}
