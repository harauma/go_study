package main

import "fmt"

//構造体
//メソッド
//構造体とスライス

type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	//フィールド名が同じ場合は型名省略できる
	Point
}

type Person struct {
	Name string
}

type Persons struct {
	Persons []*Person
}

func Update(p Point) {
	p.A = 100
	p.B = "update"
	p.C = 2.14
}

func Update2(p *Point) {
	p.A = 100
	p.B = "update"
	p.C = 2.14
}

func Set(p *Point, i int) {
	p.A = i
}

//structメソッドのレシーバーは基本的にポインタ型
func (p *Point) Set1(i int) {
	p.A = i
}

func (p Point) Set2(i int) {
	p.A = i
}

/*
コンストラクタ
New + 構造体名
基本はポインタ型を返す
*/
func NewPoint(a int, b string, c float64) *Point {
	return &Point{a, b, c}
}

func main() {
	var p Point
	fmt.Println(p)

	p2 := Point{A: 1, B: "Go", C: 1.2}
	fmt.Println(p2)
	fmt.Println(p2.A, p2.B, p2.C)

	p2.A = 10
	fmt.Println(p2.A)

	p3 := Point{1, "Python", 3.14}
	fmt.Println(p3)

	p4 := Point{A: 2}
	fmt.Println(p4)

	p5 := Point{}
	Update(p5)
	fmt.Println(p5)

	//推奨のやり方
	p6 := &Point{}
	Update2(p6)
	fmt.Println(p6)

	Update2(&p5)
	fmt.Println(p5)

	//推奨ではないがp6と同様な宣言方法
	p7 := new(Point)
	Update2(p7)
	fmt.Println(p7)

	p8 := &Point{A: 1}
	Set(p8, 2)
	fmt.Println(p8)

	p8.Set1(3)
	fmt.Println(p8)

	p9 := Point{}
	p9.Set2(100)
	fmt.Println(p9)

	p9.Set1(200)
	fmt.Println(p9)

	bp := BigPoint{}
	fmt.Println(bp)
	bp.Point.A = 100
	bp.Point.B = "Apple"
	bp.Point.C = 2.8
	fmt.Println(bp)

	bp2 := BigPoint{
		Point: Point{
			A: 100,
			B: "Banana",
			C: 3.5,
		},
	}
	fmt.Println(bp2)

	bp2.Point.Set1(2000)
	fmt.Println(bp2)

	//コンストラクタ
	p10 := Point{1, "A", 1.1}
	p11 := NewPoint(2, "B", 2.2)
	fmt.Println(p10)
	fmt.Println(p11)

	ps := make([]Person, 5)
	fmt.Println(ps)

	ps[0].Name = "Mike"
	ps[1].Name = "Jon"
	ps[2].Name = "Nina"
	ps[3].Name = "Joe"
	ps[4].Name = "Nancy"

	fmt.Println(ps)

	ps2 := Person{"Mike"}
	ps3 := Person{"Jon"}
	ps4 := Person{"Nina"}
	ps5 := Person{"Joe"}
	ps6 := Person{"Nancy"}

	ps7 := Persons{}
	ps7.Persons = append(ps7.Persons, &ps2, &ps3, &ps4, &ps5, &ps6)

	for _, p := range ps7.Persons {
		fmt.Println(p)
	}

}
