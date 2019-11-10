package main

import (
	"fmt"
	"time"
)

/* ------------------------- 无缓冲区channel ------------------------------- */

/* -------------- 相加 -------------- */
func sum(a []int, ch chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println("我将 10s 后计算完成再将数据 push 到通道 ch")
	time.Sleep(time.Second * 10)
	ch <- sum // 将 sum 传回，可以理解为 push 到 ch 通道
	fmt.Println("消息push完毕", sum)
}

/* -------------- main函数 -------------- */
func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int)
	// 这两个顺序不一定，谁先返回先打印谁
	go sum(a[:len(a)/2], ch) // 开启一个协程 并用 ch 作为通信通道
	go sum(a[len(a)/2:], ch)

	fmt.Println("ch 中没有数据我将阻塞起来等数据")
	x, y := <-ch, <-ch // 接收 ch 通道中的值
	fmt.Println(x, y, x+y)

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	// 17 -5 12 或 -5 17 12

}
