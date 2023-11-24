package main

import (
	"fmt"
	"time"
)

// forループでインクリメントされるiを送信し、recieverで受信し表示する
// 上記処理をゴルーチンで並列処理する

func reciever(c chan int) {
	// チャネルを受信して出力する
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// ゴルーチンで並列処理
	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		// 各チャネルにiを送信
		ch1 <- i
		ch2 <- i
		time.Sleep(500 * time.Millisecond)
		i++
	}
}
