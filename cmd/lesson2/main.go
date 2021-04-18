package main

import "fmt"

//関数
//クロージャーの実装

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

//ジェネレータの実装
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))

	i := integers()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	
}
