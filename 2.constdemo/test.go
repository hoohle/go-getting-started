package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	const ( //常量还可以用做枚举
		Unknow = 0
		Female = 1
		Male   = 2
	)

	const (
		A = "agoihadigohdaiosghaoidshgioadshgoikdsghiaods"
		B = len(A)
		C = unsafe.Sizeof(A)
	)

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)

	println(Unknow, Female, Male)
	println(A, B, C)
}
