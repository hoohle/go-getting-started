package main

import "fmt"

/*Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：

    Title ：标题
    Author ： 作者
    Subject：学科
    ID：书籍ID
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	//创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	//也可以使用键值对key==>value的形式		被忽略的字段为0或空
	fmt.Println(Books{title: "GOlang", author: "www.....", subject: "golang教程"})

	//如何访问结构体成员呢？	如果要访问结构体成员，需要使用点号 . 操作符，格式为： 结构体.成员名"
	var book1 Books //声明book1为Books类型
	var Book2 Books

	/* book 1 描述 */
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407
	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.book_id)
	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	//结构体作为函数参数
	/* 打印 Book1 信息 */
	printBook(book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

}

//结构体作为函数参数
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
