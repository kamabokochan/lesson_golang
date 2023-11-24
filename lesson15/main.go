package main

import "fmt"

// ポインタ

func Double(i *int) {
	*i = *i * 2
}

func main() {
	n := 100
	fmt.Println(n)  // 100
	fmt.Println(&n) // 0xc000108018(ランダムなメモリ上のアドレス)

	// *でポイント型となり、アドレス(&n)を代入している
	var p *int = &n
	fmt.Println(p)  // nと同じアドレスが表示される
	fmt.Println(*p) // アドレスにある値を参照する（データ本体を参照）

	// 互いに参照先が同じなのでデータの上書きが可能
	// *p = 300
	// fmt.Println(n) // 300
	// n = 200
	// fmt.Println(*p) // 200

	// アドレスを渡すことで、nが書き変わる
	Double(&n)
	fmt.Println(n)
}
