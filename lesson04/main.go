package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

func Price(x int) (result int) {
	result = x * 2
	return
}

func main() {
	var result = Plus(1, 2)
	fmt.Println(result)

	price := Price(1000)
	fmt.Println(price)
}
