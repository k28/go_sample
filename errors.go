package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	/*
	   Goのプログラムはエラーの状態をerror値で表現する
	   fmtパッケージはerror interfaceがnilでない値を返すとその値を表示する
	   nilの時には成功した事を表す
	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
