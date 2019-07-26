package main

//这哪个王八羔子写的代码也太特么烧脑了
import "fmt"

// 声明一个函数类型
type cb func(int) int

func main() {
	testCallBack(1, callBack)         //把callBack作为参数传递，callback需要参数，为什么此处没写呢，因为testCallBack函数允许第二个参数是一个 func(int) int类型,
	testCallBack(2, func(x int) int { //此处不用cb类型作为第二个参数直接放入函数体作为参数，因为testcallback作为函数体的第二个参数本身就是函数，所以类型完全吻合
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	testCallBack(3, callBack2)
}

func testCallBack(x int, f cb) {
	f(x) //直接把本因作为参数传入的x 作为第二个参数（函数体）的参数，		这样做的操作的怪异点就是你testcallback的时候第二个参数的函数可以不用写参数，因为参数已经被f(x)赋值了就是地一个参数，不信请看第二个函数测试
}

func callBack(x int) int {
	fmt.Printf("我是callBack，x：%d\n", x)
	return x
}

func callBack2(x int) int {
	y := x * x * x
	fmt.Printf("我要返回数的平方%d\n", y)
	return y
}
