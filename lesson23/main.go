// Stringer
package main

import "fmt"

// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

// func (p *Point) String() string {
// 	return fmt.Sprintf("%v, %v", p.A, p.B)
// }

func main() {
	p := &Point{A: 100, B: "ABC"}
	// fmt.PrintlnはString()メソッドを優先して処理の呼び出しを行う
	fmt.Println(p)
}
