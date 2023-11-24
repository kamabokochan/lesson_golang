package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// reciever
// User型のNameを表示する
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// ポインタを受け取る
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName() // user1

	user1.SetName("A")
	// 参照型ではないため、値が更新されない
	user1.SayName() // user1

	// ポインタ型で参照を渡すことで値が更新される
	user1.SetName2("A")
	user1.SayName() // A
}
