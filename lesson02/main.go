package main

import "fmt"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	// 一括宣言
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int    // 初期値 0
	var s3 string // 初期値 空文字("")
	fmt.Println(i3, s3)

	// 値の代入
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な定義
	// 変数名 := 値

	i4 := 400
	fmt.Println(i4)

	// 関数を呼び出す
	outer()
}
