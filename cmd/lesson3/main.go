package main

import (
	"fmt"
)

//go goroutin

// func sub() {
// 	for {
// 		fmt.Println("Sub Loop")
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

//init
func init() {
	fmt.Println("init")
}

func main() {
	// go sub()
	// go sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Microsecond)
	// }

	//スライス

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl)
	fmt.Println(n, sl2)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"

	s, ok := m4[1]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	delete(m4, 2)

	for k, v := range m {
		fmt.Println(k, v)
	}

	var chl chan int

	//受信専用
	// var ch2 <-chan int

	//送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))

	i := <-ch3
	fmt.Println(i)

	i2 := <-ch3
	fmt.Println(i2)

	fmt.Println("len", len(ch3))

	ch3 <- 1
	fmt.Println("len", len())
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6

}
