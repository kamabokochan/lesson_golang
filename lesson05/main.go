package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// i: index
	// v: 要素
	for i, v := range arr {
		fmt.Println(arr[i], v)
	}

	sl := []string{"Python", "PHP"}
	for i, v := range sl {
		fmt.Println(sl[i], v)
	}

	// map
	// apple 100
	// banana 200
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
