package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)        // [0,0,0]
	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	fmt.Println(len(arr4)) // len関数で配列の要素数を取得

	const (
		c0 = iota
		c1
		c2
	)

	fmt.Println(c0, c1, c2)
}
