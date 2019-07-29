package main

/*如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
*/
import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int
	a = 300

	//指针ptr地址
	ptr = &a
	pptr = &ptr

	//获取pptr的值
	fmt.Printf("变量a=%d\n", a)
	fmt.Printf("指针变量 *ptr=%d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
