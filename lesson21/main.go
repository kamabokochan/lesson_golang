// interface
package main

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Number string
	Model  string
}

// レシーバの引数が異なれば、同じ関数名を宣言可能

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString()) // Person型もCar型もToStringを持っているため、型が異なっても処理が可能
	}
}
