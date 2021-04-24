package main

import (
	"fmt"
	. "fmt"
	f "fmt"
	"golang_udemy/cmd/lesson8/foo"
)

//スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	//sの値が表示される
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	f.Println(foo.ReturnMin())
	Println(foo.ReturnMin())

	fmt.Println(appName())

	fmt.Println(Do("AAA"))
}
