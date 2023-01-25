package main

import (
	"fmt"
)

func main() {
	var i int = 100
	/*
		最大値・最小値
			int8
			int16
			int32
			int64
	*/
	var i2 int64 = 200

	fmt.Println(i + 50)

	// int型が異なるものの足し算はできない
	// fmt.Println(i + i2)

	// 書式を印刷する
	fmt.Println("%T\n", i2)

	fmt.Println("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}
