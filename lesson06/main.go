// 型アサーション

package main

import "fmt"

func anything(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}

func main() {
	anything(3)
	anything("string")
}
