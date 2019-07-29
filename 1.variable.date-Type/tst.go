package main

/*
tip1:所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：
tip2:当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝：
 tip3:可以通过 &i 来获取变量 i 的内存地址

tip4:更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
tip5:一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。
ps:这个内存地址为称之为指针，这个指针实际上也被存在另外的某一个字中。
同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；(数据结构知识)

tip6: 当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。
tip7:如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。

*/
import "fmt"

//以下几种类型为nil
// var a *int
// var a []int
// var a map[string]int
// var a chan int
// var a func(string) int
// var a error // error 是接口

var ( //这种因式分解关键字的写法一般用于声明全局变量
	pub1 bool
	pub2 string
	pub3 int
	pub4 int
)

//这种全局变量允许声明但不使用，局部的不行

func main() {
	var a = "Runboo"              //此处go会自动推断类型，所以无需声明
	a2, a22 := "Runboo", "Google" //此处直接省略了var	因为:=是一个声明语句	ps:这种写法只能出现在函数体中（不带声明格式的）
	fmt.Println(a)
	fmt.Println(a2, a22)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var a1 = "RUNOOB"
	fmt.Println(a1)
	var b1 int
	fmt.Println(b1)
	var c1 bool
	fmt.Println(c1)

	var i int
	var f float64
	var j bool
	var k string
	fmt.Printf("%v	%v	%v	%q\n", i, f, j, k)

	//go可以根据值自行判断变量类型
	var d = true
	fmt.Print(d, "\n")
	fmt.Println(d)
	//上面的两句是等效的
	fmt.Print(d)
	fmt.Print(d)

	//多变量声明，非全局变量
	// var t1,t2,t3	type
	t1, t2, t3 := 1, 2, 3 //出现在：=左侧的变量不应该是已经被声明过的，否则会导致编译错误
	fmt.Println("\n", t1, t2, t3)
	var T1, T2, T3 = 4, 5, 6 //与python类似，不需要显示声明类型，自动推断
	fmt.Println(T1, T2, T3)

	//在函数体之前用因式分解关键字的写法声明的全局变量 pub1 pub2 pub3
	pub1 = true
	pub2 = "infsoif"
	pub3 = 456
	println(pub1, pub2, pub3)

	//同一行赋值
	pub1, pub2, pub3 = false, "hhhhh", 123456
	println(pub1, pub2, pub3, pub4)

	//简单的进行交换数据
	pub3, pub4 = pub4, pub3
	println(pub3, pub4)

	//因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	//_实际上是一个只写变量
	// value1, value2 := 2, 3
	// _, value2 = value1, value2

}
