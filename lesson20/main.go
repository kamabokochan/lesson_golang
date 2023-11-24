// 構造体 map
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}

	fmt.Println(m)

	m2 := make(map[int]User)
	m2[1] = User{Name: "user3", Age: 30}
	fmt.Println(m2)

	for _, v := range m {
		fmt.Println(v)
	}
}
