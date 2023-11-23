// map

package main

import "fmt"

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m) // map[A:100 B:200]

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2) // map[A:100 B:200]

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3) // map[1:A 2:B]

	// make関数
	m4 := make(map[int]string)
	fmt.Println(m4) // map[]

	m4[1] = "JAPAN"
	m4[2] = "USA"

	fmt.Println(m4) // map[1:JAPAN 2:USA]

	// 値の取り出し
	fmt.Println(m["A"]) // 100

	// 第二引数で値の有無を取得可能
	_, ok := m4[3]
	// m4の3番目に要素がないため、okはfalseになる
	if !ok {
		fmt.Println("error")
	}

	m4[3] = "CHINA"
	fmt.Println(m4) // map[1:JAPAN 2:USA 3:CHINA]

	// 値の削除
	delete(m4, 3)
	fmt.Println(m4) // map[1:JAPAN 2:USA]

	// mapの長さを取得
	fmt.Println(len(m4)) // 2
}
