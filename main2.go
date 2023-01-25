package main

import (
	"fmt"
)

// 浮動小数点型
func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Println("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Println("%T\n", fl32)

	fmt.Println("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
