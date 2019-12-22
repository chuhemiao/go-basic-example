package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//main()本身也是运行了一个goroutine。

	messages := make(chan string) //声明了一个阻塞式的无缓冲的通道

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	// 一个基本的goroutine
	go say("world")
	// 一个函数
	say("hello")

	// 使用 sync.Mutex 互斥锁类型 lock unlock 进行通信

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

}

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
