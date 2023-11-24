// 構造体のスライス
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// ポイント型のUserのスライス型で宣言
type Users []*User

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4) // 可変長引数も可能

	// fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u) // 実体で表示
	}
}
