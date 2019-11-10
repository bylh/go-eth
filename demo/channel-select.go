package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select { // ch/ch1 通道都没有数据的时候 select 会阻塞起来等待数据，两个同时有数据时候随机选一个
			case v := <-ch:
				fmt.Println("v:", v)
			case a := <-ch1:
				fmt.Println("a:", a)
			case <-time.After(2 * time.Second): // 如果前面的所有通道都阻塞了2s就执行这里
				fmt.Println("time out")
				o <- true
				break
			}
		}
	}()
	// 模拟向通道写数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i + 1
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i + 1
		}
	}()
	// 会阻塞起来等待 o 中的数据
	<-o
}
