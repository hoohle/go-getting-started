package main

import "fmt"

func main() {
	const (
		a = iota //0
		b        //1
		c        //2

		d = "ha" //独立值，iota +=1
		e        //"ha"	iota+=1

		f = 100 //iota+=1
		g       //100	iota+=1

		h = iota //7恢复计数
		i        //8
	)

	//在来个有趣的iota实例
	const (
		j0 = 1 << iota
		j  = 3 << iota
		k
		l
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println("j0=", j0)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
