// panic recover

package main

import "fmt"

func main() {
	defer func() {
		// panicで発生したランタイムエラーから復帰する機能
		// deferの中でしか使われないため、セットになる
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// ランタイムエラーが発生し、処理が強制終了する
	// panic はアプリケーション内で続行不可能な致命的エラーが発生した場合に投げられる
	panic("runtime error")
}
