package main

import (
	"fmt"
	"golang_udemy/cmd/lesson9_test/alib"
)

//テスト

func IsOne(i int) bool {
	return i == 1
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
