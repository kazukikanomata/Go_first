package main

import (
	"fmt"
)

func main() {
	var s string = "hello Golang"
	fmt.Println(s)
	fmt.Println("%T\n", s)

	var si string = "300"

	fmt.Println(si)
	fmt.Println("%T\n", si)

	fmt.Println(`test
	test
	test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))

}
