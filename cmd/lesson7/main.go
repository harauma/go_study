package main

import "fmt"

//interface
//最もポピュラーな使い方。異なる型に共通の性質を付与する。
//カスタムエラー
//fmt.Stringer

/*
type error interface {
	Error() string
}
*/

/*
type Stringer interface {
	String() string
}
*/

type Stringfy interface {
	ToString() string
}

type Point struct {
	A int
	B string
}

type MyError struct {
	Message string
	ErrCode int
}

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Number string
	Model  string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	p := &Point{100, "ABC"}
	fmt.Println(p)
}
