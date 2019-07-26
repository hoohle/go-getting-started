package main

import (
	"fmt"

	demoparameter "./parameter_demo"
)

/*通过函数划分不同功能，逻辑上每个函数执行的是指定的任务
函数声明告诉了编译器函数的名称，返回类型，和参数
标准库提供了多种可用的内置函数，例如len()函数接受不同类型参数并返回该类型长度。
如果是字符串则返回字符串长度，如果传入数组，则返回数组中包含元素的个数
*/

/*Go 语言函数定义格式如下：

func function_name( [parameter list] ) [return_types] {
	函数体								参数列表					返回类型
 }*/
/*function return the max number between two numbers*/
func max(num1, num2 int) int {
	var result int
	//declare local area variable
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//GO函数可以返回多个值，例如：
func swap(a, b string) (string, string) {
	return b, a
}

/*当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

调用函数，向函数传递参数，并返回值，*/
func main() {
	//declare local area variable
	var a = 300
	var b = 301
	var ret int
	//declare function and return the max number
	ret = max(a, b)
	fmt.Printf("最大值是：%d\n", ret)

	x, y := swap("Google", "Runoob") //此处不能再次使用a和b因为已被声明，且类型也不一样
	//换值函数
	fmt.Println(x, y)
	fmt.Println(swap("world", "Hello")) //直接打印函数，谁说的一定要有接收值，函数体本身即变量

	//GO语言函数作为实参
	/*声明函数变量*/
	var typegetSquareRoot float64
	typegetSquareRoot = demoparameter.Realvalue9(90)
	getSquareRoot2 := demoparameter.Realvalue9(100)
	/*use the function*/
	fmt.Println(typegetSquareRoot)
	fmt.Println(getSquareRoot2)
	fmt.Println(demoparameter.Realvalue9(9))

	A := 100
	B := 200
	fmt.Printf("交换前，a 的值 : %d\n", A)
	fmt.Printf("交换前，b 的值 : %d\n", B)
	/* 调用 swap2() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	demoparameter.Swap2(&A, &B)
	fmt.Printf("交换后，a的值: %d\n", A)
	fmt.Printf("交换后，b的值%d\n", B)
}
