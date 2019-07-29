package main

import "fmt"

/*二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：
var arrayName [ x ][ y ] variable_type
			数组名                               类型
*/
//初始化二维数组
func main() {
	//数组-5行2列
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	//输出数组元素
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
