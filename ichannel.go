package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

// 关闭通道

func fibonaClose(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 通道的同步

// 这个worker函数将以协程的方式运行
// 通道`done`被用来通知另外一个协程这个worker函数已经执行完成

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	// 用来通知工作完成，向通道发送一个数据，表示worker函数已经执行完成
	done <- true
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)

	// 带有缓冲的channel

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 关闭一个通道
	cclose := make(chan int, 10)
	go fibonaClose(cap(cclose), cclose)
	for i := range cclose {
		fmt.Println(i)
	}
	// 通道的同步

	// 使用协程来调用worker函数，同时将通道`done`传递给协程
	// 以使得协程可以通知别的协程自己已经执行完成
	done := make(chan bool, 1)
	go worker(done)
	// 一直阻塞，直到从worker所在协程获得一个worker执行完成的数据
	<-done

}
