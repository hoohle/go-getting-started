package main

import "fmt"

/*
Go 语言提供了数组类型的数据结构。
数组是具有相同唯一类型的一组已编号且 长度固定 的数据项序列，这种类型可以是任意的原始类型，如整形、字符串或者自定义类型
相对于去声明num0、num1、。。。，num99的变量，使用数组形式更加方便易于扩展
数组元素可以通过索引位置来读取或者修改
*/

//1.声明数组
//2.初始化数组		类似于先声明再定义？

//以下演示了数组完整操作（声明、赋值、访问）的实例：
func main() {
	var n [10]int
	var j int

	//为数组n初始化元素
	for i := 0; i < 10; i++ {
		n[i] = i + 100 //设置元素为i+100
	}

	//输出每个数组元素的值		遍历？
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d]=%d\n", j, n[j])
	}

	//杨辉三角函数
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = GetYanghuiTriangleNextLine(nums)
		fmt.Println(nums)
	}

	//第二个杨辉三角
	yanghuisanjiao(12)
}

//来一段高级点的示例，用数组打印杨辉三角
func GetYanghuiTriangleNextLine(inArr []int) []int {
	var out []int
	var i int
	arrlen := len(inArr)
	out = append(out, 1)
	if 0 == arrlen {
		return out
	}
	for i = 0; i < arrlen-1; i++ {
		out = append(out, inArr[i]+inArr[i+1])
	}
	out = append(out, 1)
	return out
}

func yanghuisanjiao(rows int) {
	for i := 0; i < rows; i++ {
		number := 1
		for k := 0; k < rows-i; k++ {
			fmt.Print("  ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%5d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
