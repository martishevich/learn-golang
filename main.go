package main

import (
	"fmt"
	"strconv"
)

func myStrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func myStrToInt2(s string) (int, error) {
	var number int
	_, err := fmt.Sscanf(s, "%d", &number)
	return number, err
}

func main() {
	res, err := myStrToInt("3")
	if err != nil {
		fmt.Println(err)
		panic("")
	}
	fmt.Println(res)


	res2, err2 := myStrToInt2("3")
	if err2 != nil {
		fmt.Println(err2)
		panic("")
	}
	fmt.Println(res2)
}