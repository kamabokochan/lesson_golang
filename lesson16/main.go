package main

import "fmt"

type User struct {
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Age = 15
	user.Name = "A"
}

func UpdateUser2(user *User) {
	user.Age = 15
	user.Name = "A"
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Age = 10
	user1.Name = "user1"
	fmt.Println(user1)

	user2 := User{}
	user2.Age = 20
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	// ポインタ
	user4 := new(User)
	fmt.Println(user4) // &{ 0}
	user5 := &User{}
	fmt.Println(user5) // &{ 0}

	// 参照が渡らないため値は更新されない
	UpdateUser(user1)
	fmt.Println(user1) // {user1 10}

	// ポイントで参照が渡るため値が更新される
	UpdateUser2(user5)
	fmt.Println(user5)  // &{A 15} (ポイント型)
	fmt.Println(*user5) // {A 15} (データ本体)

}
