package main

import "fmt"

func main() {
	// var count,c int   //定义变量不使用也会报错
	var count int
	var flag bool
	count = 1
	//while(count<100) {    //go没有while
	for count < 100 {
		count++
		flag = true
		//注意tmp变量  :=
		for tmp := 2; tmp < count; tmp++ {
			if count%tmp == 0 {
				flag = false
			}
		}

		// 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
		if flag == true {
			fmt.Println(count, "素数")
		} else {
			continue
		}

	}
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)

	// var A int
	// var B int
	// fmt.Printf("请输入密码：   \n")
	// fmt.Scan(&A)
	// if A == 5211314 {
	// 	fmt.Printf("请再次输入密码：")
	// 	fmt.Scan(&B)
	// 	if B == 5211314 {
	// 		fmt.Printf("密码正确，门锁已打开")
	// 	} else {
	// 		fmt.Printf("非法入侵，已自动报警")
	// 	}
	// } else {
	// 	fmt.Printf("非法入侵，已自动报警")
	// }

	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	//Type Switch 语法格式如下
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		// fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		// fallthrough
	default:
		fmt.Println("6、默认 case")
	}


	/*
	以下描述了 select 语句的语法：

    每个 case 都必须是一个通信
    所有 channel 表达式都会被求值
    所有被发送的表达式都会被求值
    如果任意某个通信可以进行，它就执行，其他被忽略。
    如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
    否则：
        如果有 default 子句，则执行该语句。
        如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("sent", i2, "to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("c3 is closed\n")
	}
}
