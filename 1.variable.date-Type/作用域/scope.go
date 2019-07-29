/*
作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

    函数内定义的变量称为局部变量
    函数外定义的变量称为全局变量
    函数定义中的变量称为形式参数

接下来让我们具体了解局部变量、全局变量和形式参数。
*/
package main

import "fmt"

//声明全局变量
var g int
var c int = 0

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}

func main() {
	//声明局部变量
	var a, b, c int
	a = 10
	b = 20
	c = a + b
	fmt.Printf("结果:a=%d,b=%d and c=%d\n", a, b, c)

	g = a + b
	fmt.Printf("结果:a = %d, b = %d and g = %d\n", a, b, g)

	//Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
	fmt.Printf("c的值是全局的0呢还是局部的30： g = %d\n", c)

	/*
		形式参数

	形式参数会作为函数的局部变量来使用。实例如下：*/
	fmt.Printf("main()函数中 a = %d\n", a)
	C := sum(a, b)
	fmt.Printf("main()函数中 C = %d\n", C)

}
