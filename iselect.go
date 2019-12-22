package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1) //定义两个有缓冲通道，容量为1
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1) //每隔1秒发送数据
		c1 <- "name: chuhemiao"
	}()

	go func() {
		time.Sleep(time.Second * 2) //每隔1秒发送数据,如果时间超过了接收并写入的时候，select会走超时入口
		c2 <- "age: 18"
	}()

	for i := 0; i < 2; i++ { //使用select来等待这两个通道的值，然后输出
		tm := time.NewTimer(time.Second * 5)
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-tm.C:
			fmt.Println("send data timeout!")
		}
	}

	// select默认选择

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("默认情况....")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
