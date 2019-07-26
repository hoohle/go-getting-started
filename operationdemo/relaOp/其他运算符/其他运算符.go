package main

func main() {
	var a int = 4
	b := 5
	var ptr *int
	ptr = &a

	println("a的值为", a)     // 4
	println("*ptr为", *ptr) // 4	指向指针的指针	此处是4字节的指针
	println("ptr为", ptr)   // 824633794744
	ptr = &b
	println("*ptr为", *ptr) // 4	指向指针的指针	此处是4字节的指针
	println("ptr为", ptr)   // 824633794744

	// pttr = **a
	// println("pttr为", pttr) // 824633794744
}
