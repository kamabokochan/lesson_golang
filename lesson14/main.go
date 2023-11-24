package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "A"

	// switchと違い、どのcase式が実行されるかはランダム
	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	default:
		fmt.Println("どちらでもない")
	}

	// 例2
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3   // 2: ch3が受信しiに代入
			ch4 <- i * 2 // 3: ch4に送信
		}
	}()

	go func() {
		for {
			i2 := <-ch4   // 4: ch4が受信し、iに代入
			ch5 <- i2 - 1 // 5 ch5に送信
		}
	}()

	n := 0

	// チャネルへの送受信は実行可能かを判断して、可能であれば実行される。
	// 送信の場合は「キャパシティ（バッファ）がいっぱいになっていないか」、
	// 受信の場合は、「チャネルに値が送信されたか、チャネルが閉じられたか」を確認し、処理を進める準備が出来ていれば実行される。
L:
	for {
		select {
		case ch3 <- n: // 1: ch3に送信した時に実行
			n++
		case i3 := <-ch5: // ch5が受信し、i3に代入
			fmt.Println("recieve", i3)
		default:
			if n > 100 { // nが100より大きくなったらbreakで終了
				break L
			}
		}
		// if n > 100 { // nが100より大きくなったらbreakで終了
		// 	break
		// }
	}
}

// select文の仕様
// select {
// case <-ch1:
// 	// ch1から受信したときに実行される処理
// case v := <-c2:
// 	// ch2から受信したときに実行される処理
// 	// 変数に値を入れることもできる
// case ch3 <- y:
// 	// ch3に送信したときに実行される処理
// default:
// 	// どのcaseも準備できていないときに実行される処理（省略可能）
// }
