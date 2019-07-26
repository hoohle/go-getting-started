package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)

	}
}

// ch := make (chan int)
/*
 注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须在接收端相应的接收数据。

以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把sum发送到通道c
}

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int)
	c := make(chan int, 100)
	// 通道可以设置缓冲区，通过设置make的第二个参数指定缓冲区的大小
	/*缓冲区解释
		带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

	不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

	注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
	*/

	//以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c       //从通道c中接收
	fmt.Println(x, y, x+y) //菜鸟go是-5	17 	12

	//演示缓冲区
	channel := make(chan int, 2)
	//因为channel是带缓冲区的通道，我们可以同时发送两个数据		而不用立刻需要去同步读取数据
	channel <- 1
	channel <- 2
	//获取这两个数据
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
