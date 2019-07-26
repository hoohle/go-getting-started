package main

import (
	"fmt"

	relationship "./relaOp"
)

/*
+	相加
-	相减
*	相乘
/	相除
%	求余
++	自增
--	自减
*/

func main() {
	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("line1-c cost value is:%d\n", c)
	//31
	c = a - b
	fmt.Printf("line2-c cost:%d\n", c)
	//11
	c = a * b
	fmt.Printf("a*b=%d\n", c)
	//210
	c = a / b
	fmt.Printf("a/b=%d\n", c)
	//2
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	//?1
	a++
	fmt.Printf("a=	%d\n", a)
	//22
	// a = 21
	a--
	fmt.Printf("a=	%d\n", a)

	fmt.Printf(relationship.Isequal(4, 3))
	fmt.Println()
	fmt.Printf(relationship.Logic(true, true))
	// fmt.Printf(relationship.Bit2(60, 13))
	fmt.Println(relationship.Ass(23, 2))

}
