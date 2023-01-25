package main

// https://go.dev/doc/tutorial/create-module

// formatぱケージ
// go は一個のpackage

import (
	"fmt"
	"time"
)

// hello world
func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())
}
