package main

import "fmt"

const MAX int = 3

func main() {
	// a := []int{10, 100, 200}
	// var i int

	// for i = 0; i < MAX; i++ {
	// 	fmt.Printf("a[%d]=%d\n", i, a[i])
	// }

	//有一种情况，我们可能需要保存数组，这样我们就需要使用到指针。
	//以下声明整数型指针数组
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
