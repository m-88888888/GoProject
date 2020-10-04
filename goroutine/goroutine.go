package main

import (
	"fmt"
	"time"
)

func hoge() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(100 * time.Millisecond)
	}
}

func fuga() {
	go hoge()
	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(200 * time.Millisecond)
	}
}

func channel(chA chan<- string) {
	time.Sleep(3 * time.Second)
	chA <- "Finished"
}

func main() {
	fuga()

	chA := make(chan string)
	defer close(chA)
	fmt.Println("ゴルーチン実行開始")
	go channel(chA)
	msg := <-chA
	fmt.Println(msg)
}
