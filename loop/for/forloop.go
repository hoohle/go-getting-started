package main

import "fmt"

func main() {
	/*for的循环有三种形式
		第一种：只有这一种有分号
		...
	第三种数组的遍历

	for循环嵌套
	*/

	//for init; condition; post { }				参数解释
	//    init： 一般为赋值表达式，给控制变量赋初值；
	//  condition： 关系表达式或逻辑表达式，循环控制条件；
	//    post： 一般为赋值表达式，给控制变量增量或减量。
	b := 15
	var a int
	numbers := []int{1, 2, 3, 5}
	//第一种格式for循环
	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为：%d\n", a)
	}
	//第二种for循环
	for a < b { //14
		a++ //15
		fmt.Printf("a的值为：%d\n", a)
	}
	//第三种for循环	遍历数组式子
	for num, x := range numbers {
		fmt.Printf("第%d位x的值=%d\n", num, x)
	}

	//for循环嵌套的演示
	var i, j int

	for i = 2; i < 10; i++ {
		fmt.Printf("开始外层循环此时的i=%d	j=%d\n", i, j)
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				fmt.Printf("当前的i=%d不是素数	j值是%d\n", i, j)
				break // 如果发现因子，则不是素数	如果不break则执行下方的if语句
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}

}
