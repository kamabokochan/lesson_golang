// スライス

package main

import "fmt"

func main() {
	var sl []int
	fmt.Println(sl) // []

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // [100 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3) // [A B]

	sl4 := make([]int, 5) // 型, 要素数
	fmt.Println(sl4)      // [0 0 0 0 0]

	sl2[0] = 1000    // 値の代入
	fmt.Println(sl2) // [1000 200]

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5) // [1 2 3 4 5]

	// 値の取り出し(0番目を取り出す)
	fmt.Println(sl5[0]) // 1

	// 2番から4番を指定
	fmt.Println(sl5[2:4]) // [3 4]

	// 2番目までを指定（2番目は含めない）
	fmt.Println(sl5[:2]) // [1 2]

	// データの追加(append)
	sl2 = append(sl2, 300, 400, 500)
	fmt.Println(sl2)

	// len関数で要素数を取得
	sl6 := make([]int, 5)
	fmt.Println(len(sl6))

	// capでキャパシティを取得
	fmt.Println(cap(sl6))

	// 第二引数: 要素数
	// 第三引数: 容量
	sl7 := make([]int, 5, 10)
	fmt.Println(len(sl7)) // 5
	fmt.Println(cap(sl7)) // 10

	// copy
	sl8 := []int{100, 200}
	sl8b := sl8
	sl8b[0] = 1000
	// 参照型は、代入された変数（sl8b）の値を更新しても、
	// 参照元と同じメモリのアドレスを参照してしまうため、参照元の値も更新されてしまう
	fmt.Println(sl8) // [1000, 200]

	sl9 := []int{1, 2, 3}
	sl9b := make([]int, 5, 10)

	fmt.Println(sl9b) // [0,0,0,0,0]

	// n: コピーした要素数
	n := copy(sl9b, sl9)
	fmt.Println(n, sl9b) // 3 [1,2,3,0,0]

	// slice for
	var slice = []string{"Golang", "Ruby", "Javascript", "Python"}
	for index, value := range slice {
		fmt.Println(index, value)
		//=> 0 Golang
		//=> 1 Ruby
		//=> 2 Javascript
		//=> 3 Python
	}
}
