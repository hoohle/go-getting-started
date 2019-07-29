//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。
//ps:匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
package main

import (
	"fmt"
)

func getSequence() func() int { //返回值竟然是一个无参的匿名函数，
	i := 0
	return func() int { //函数本身即返回值
		i++
		return i
	}
}

// 闭包使用方法带参数
func add(x1, x2 int) func(x3 int, x4 int) (int, int, int) { //第二个匿名函数
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	//带参数的闭包函数调用:
	add_func := add(1, 2)
	fmt.Println(add_func(1, 1))
	fmt.Println(add_func(0, 0))
	fmt.Println(add_func(2, 2))

}
