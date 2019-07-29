/*Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务。
接下来让我们来一步步学习 Go 语言指针。
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。*/
package main

import "fmt"

func main() {

	var a int = 100
	fmt.Printf("变量的地址:%x\n", &a)

	//什么是指针，一个指针变量指向了一个值的内存地址，类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
	//var var_name *var-type

	/*如何使用指针

	  指针使用流程：

	      定义指针变量。
	      为指针变量赋值。
	      访问指针变量中指向地址的值。
	*/
	var b int = 20 //defined real variable
	var ip *int    //defined pointer variable
	ip = &b        //指针变量的存储地址
	fmt.Printf("b变量的地址是：%x\n", &b)

	//指针变量的存储地址
	fmt.Printf("ip变量储存的指针地址是：%x\n", ip)

	//使用指针访问值
	fmt.Printf("*ip变量的值:%d\n", *ip)

	/*空指针:
		当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	nil 指针也称为空指针。
	nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	一个指针变量通常缩写为 ptr。
	*/
	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
	fmt.Printf("ptr变量储存的指针地址是：%x\n", ptr)
	fmt.Printf("ptr变量的地址是：%x\n", &ptr)
	// fmt.Printf("*ptr变量指向的值是:%d\n", *ptr)
	if ptr != nil {
		fmt.Println("不是空指针")
	} else {
		fmt.Println("是空指针")
	}

}
