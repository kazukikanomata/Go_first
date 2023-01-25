package main

import (
	"fmt"
)

// https://go.dev/doc/tutorial/create-module

// formatぱケージ
// go は一個のpackage
// 明示的な定義
// var i int = 100

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {

	// var 変数名　型　= 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義

	// 変数名　:=　値
	// 最定義ができないよ
	i4 := 400
	fmt.Println(i4)

	// i4 := 450
	// fmt.Println(i4)

	// i4 := 500
	// fmt.Println(i4)

	// i4 = "hello"
	// fmt.Println(i4)

	outer()

	// fmt.Println(s4)

	// 定義されていない変数は使われていないの判定がある、
	// var s5 string = "Not Use"
}
