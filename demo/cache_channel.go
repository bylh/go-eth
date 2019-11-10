package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		// 延迟2s发送
		time.Sleep(time.Second * 2)
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	// 信道缓存为10
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// 循环遍历信道，取得结果
	for i := range c {
		fmt.Println(i)
	}
}
