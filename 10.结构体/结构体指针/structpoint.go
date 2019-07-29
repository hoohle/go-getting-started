package main

import "fmt"

type Books struct { //struct类似于JAVA 中的类，创建一个对象，总结出他的属性
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book *Books) { //结构体指针参数
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {
	var book1 Books //用属性具现化对象
	var book2 Books

	//book1描述
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407

	/* book 2 描述 */
	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&book1)

	/* 打印 Book2 信息 */
	printBook(&book2)

	var c Rect
	c.width = 100
	c.height = 20
	fmt.Println(c.Area())
}

//来试试golang的OO 方法
type Rect struct { //定义矩形类
	x, y          float64 //类型只包含 属性
	width, height float64
}

func (r *Rect) Area() float64 { //为Rect类型绑定Area方法，*Rect为指针引用可以修改传入参数的值
	return r.width * r.height
}
