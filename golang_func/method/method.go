/*Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。语法格式如下：*/
package main

import (
	"fmt"
)

//定义结构体
type Circle struct { //Circle类型
	radius    float64 //类中的属性
	diameter  float64
	realvalue float64 //真正的值
}

/*Go 没有面向对象，而我们知道常见的 Java。

C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针，而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别。


在 C++ 中是这样的:


class Circle {
  public:
    float getArea() {
       return 3.14 * radius * radius;
    }
  private:
    float radius;
}

// 其中 getArea 经过编译器处理大致变为
float getArea(Circle *const c) {
  ...
}*/
//而在GO中是这么写的

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 { //所谓接受者想必就是Circle类型果然被函数所接受
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius //pi r^2计算元的面积
}

func (c Circle) getperimeter() float64 {
	pid := 3.14 * c.diameter
	pir := 2 * 3.14 * c.radius
	if pid == pir {
		c.realvalue = pid
	} else {
		c.realvalue = 0
		fmt.Printf("两个值算出来竟然不想等")
	}
	return c.realvalue
}

func main() {
	var c1 Circle
	c1.radius = 10.00000
	c1.diameter = 20.0000
	fmt.Println("圆的面积=", c1.getArea())
	fmt.Println("圆的周长=", c1.getperimeter())
}

/*第一个简单的例子：编写一个驾驶汽车的方法

面向过程的程序设计：
编写一个方法，void drivecar();
面向对象的程序设计：
将一辆汽车看成一个对象，将所有汽车对象的共性抽取出来，设计一个类Car，类中有一个方法void drive()，用Car这个类实例化一个具体的对象car，调用：car.drive()。
Java的世间万物皆对象，对象里有类，类里有方法，用汽车这个类实例化一个具体对象 调用car.driver


第二个简单的例子：求一个长方形的周长和面积。
面向过程的程序设计方式：
1、确定长方形周长和面积的算法。
2、编写两个方法（函数）分别计算长方形的周长和面积。
3、求周长的方法（函数）和求面积的方法（函数）需要两个参数，分别是长方形的长和宽。

面向对象的程序设计方式：//我把这种思想称为JAVA思想
1、一个长方形可以看成一个长方形对象。
2、一个长方形对象有两个状态（长和宽）和两个行为（求周长和求面积）。
3、将所有长方形的共性抽取出来，设计一个长方形类。
4、通过长方形对象的行为，就可以求出某个具体的长方形对象的周长和面积。
*/
