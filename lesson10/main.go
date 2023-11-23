package main

import "fmt"

// 可変長引数
// 引数を無限に渡すことができる

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	// スライスを渡すことも可能
	sl := []int{1, 2, 3}
	fmt.Println(sum(sl...)) // 6
}
