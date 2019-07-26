// package demo_parameter		//错误示范不能加_符号
package demoparameter

import "math"

/*if you use parameter ,it means the variable called 形参
形参就像定义在函数体内的局部变量。
调用函数，可以通过两种方式来传递参数*/
/*
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

第一种较为常见：值传递
1.值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
个人理解就是函数体前边的是形参可以声明不赋值，然后实参作为值传递给形参，这样实参的值不会变
2。第二种	（高级方法）引用传递
引用传递是指在调用函数时将		实际参数的地址		传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
*/
// func realvalue9(x float64) float64 {//错误示范函数名不大写，死都掉不出去
func Realvalue9(x float64) float64 {
	return math.Sqrt(x)
}

/*引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：*/
func Swap2(x *int, y *int) {
	var temp int
	temp = *x //保持x地址上的值送给temp
	*x = *y   //将y值赋给x
	*y = temp //将temp值赋给y
}
