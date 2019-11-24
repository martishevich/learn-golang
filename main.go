package main

import "fmt"

func average(numbers []int) float64 {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return float64(sum) / float64(len(numbers))
}

func main() {
	result := average([]int{1,2,3,4,5,6})
	fmt.Println(result)
}