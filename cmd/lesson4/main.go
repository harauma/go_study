package main

import (
	"fmt"
	"time"
)

//channel
//チャネルとゴルーチン

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func lesson4_1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//fmt.Println(<-ch1)

	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}

func lesson4_2() {
	ch1 := make(chan int, 2)

	close(ch1)

	ch1 <- 1

	// fmt.Println(<-ch1)

	i, ok := <-ch1
	fmt.Println(i, ok)

}

func main() {
	// lesson4_1_()

}
