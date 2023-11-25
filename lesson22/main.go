// カスタムエラー
package main

import "fmt"

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

// error型にあるように、Errorメソッドを用意する
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 123}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// MyErrorを参照したい場合は、RaiseErrorから返却されるMyErrorの実体にアクセス
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
