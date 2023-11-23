package main

import "fmt"

// channel
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造

func main() {
	// 双方向チャネル
	var ch1 chan int

	// 受信専用
	// var ch2 <-chan int

	// 送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)

	// cap 容量取得
	fmt.Println(cap(ch1)) // 0

	// バッファサイズを指定
	ch2 := make(chan int, 5)
	fmt.Println(cap(ch2)) // 5

	ch2 <- 1              // ch2 に 1 を送信
	fmt.Println(len(ch2)) // 1

	ch2 <- 2
	ch2 <- 3

	// データを送信した順番に呼び出される(順番が保証される)
	i := <-ch2
	fmt.Println(i) // 1
	i2 := <-ch2
	fmt.Println(i2)    // 2
	fmt.Println(<-ch2) // 3

	// バッファ5を超える送信を行うと、deadLockとなる
	// ただし超える前に受信することで受信した分、送信可能な量が増える
}
