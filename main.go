package main

import "fmt"

func max(words []string) string {
	maxLength := len(words[0])
	maxWordId := 0
	for i, n := range words {
		if maxLength < len(n) {
			maxLength = len(n)
			maxWordId = i
		}
	}
	return words[maxWordId]
}

func reverse(numbers []int64) []int64 {
	var reversedNumbers []int64
	for _, n := range numbers {
		reversedNumbers = append([]int64{n}, reversedNumbers...)
	}
	return reversedNumbers
}

func main() {
	reversed := reverse([]int64{1,2,3,4,5,6})
	fmt.Println(reversed)

	maxWord := max([]string{"one", "two", "three"})
	fmt.Println(maxWord)

	maxWordTwo := max([]string{"one", "two"})
	fmt.Println(maxWordTwo)
}
