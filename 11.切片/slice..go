package main

/*
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
*/
import "fmt"

/*
切片的使用方法：
1.先定义切片					  你可以声明一个未指定大小的数组来定义切片：var identifier []type
切片不需要说明长度。						或使用make()函数来创建切片:				var slice1 []type = make([]type, len)		简写 slice1:=make([]type, len)
若要指定容量，其中capacity为可选参数。				make([]T, length, capacity)

2.初始化切片	ps十分灵活自己琢磨
3.len和cap函数	全看下面的实例
*/
func main() {
	numbers := make([]int, 3, 5) //建立了一个numbers切片，长度为3,容量为5
	printSlice(numbers)

	//一个切片在未初始化之前默认为nil，长度为0
	var numbers1 []int
	if numbers1 == nil {
		fmt.Printf("切片是空的\n")
		printSlice(numbers1)
	}

	//创建切片
	numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])
	numbers1 = make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
