// 並列処理 go goroutin

package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// go を先頭につけることで、並列処理されるようになる（sub関数とmain関数の処理が並行処理される）
	go sub()
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond)
	}
}
