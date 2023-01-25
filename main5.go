package main

import "fmt"

func main() {
	byteA := []byte{72, 73}

	fmt.Println(byteA)

	// 文字列がアスキーコードになる。
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

}
