package main

import "fmt"

// nill　初期値
// インターフェイスは全ての方と互換生がある
// 計算はできない
func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
}
