package main

import "fmt"

/*如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：
第一种：形参设定数组大小：void myFunction(param [10]int)
第二种：形参未设定数组大小：void myFunction(param []int)
*/
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)
	return avg
}

func main() {
	//数组长度为5
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32
	//数组作为参数传递给函数
	avg = getAverage(balance, 5)
	//输出返回的平均值
	fmt.Printf("平均值为：%f\n", avg)

	//以上实例中我们使用的形参并未设定数组大小。

	//浮点数计算输出有一定的偏差，你也可以转整型来设置精度。
	a := 1.69
	b := 1.7
	c := a * b     // 结果应该是2.873
	fmt.Println(c) // 输出的是2.8729999999999998

	//设置固定精度
	a = 1690                          // 表示1.69
	b = 1700                          // 表示1.70
	c = a * b                         // 结果应该是2873000表示 2.873
	fmt.Println(c)                    // 内部编码
	fmt.Println(float64(c) / 1000000) // 显示

}
