package main

import (
	"fmt"
	"unicode/utf8"
)

func max(words []string) string {
	if len(words) == 0 {
		return ""
	}
	maxLength := utf8.RuneCountInString(words[0])
	maxWordId := 0
	for i, n := range words {
		wordLength := utf8.RuneCountInString(n)
		if maxLength < wordLength {
			maxLength = wordLength
			maxWordId = i
		}
	}
	return words[maxWordId]
}

func reverse(numbers []int64) []int64 {
	numberLength := len(numbers)
	reversedNumbers := make([]int64, numberLength)
	for i := range numbers {
		reversedNumbers[numberLength - i - 1] = numbers[i]
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

	maxWordT := max([]string{})
	fmt.Println(maxWordT)
}
