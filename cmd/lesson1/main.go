package main

import (
	"fmt"
	"strconv"
)

//定数
const Pi = 3.14

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)

	var i2 int64 = 200
	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
}

func lesson1() {
	// fmt.Println(time.Now())
	var i int = 1000
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	//暗黙的な定義
	//変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	outer()
}

func lesson2() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr4 := [...]string{"C"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
}

func lesson3() {
	// var x interface{}
	// fmt.Println(x)

	// x = 1
	// fmt.Println(x)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)
}

func main() {
	lesson3()
}
