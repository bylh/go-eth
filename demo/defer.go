package main

import (
	"fmt"
	"os"
)

func a() int {
	defer b()
	fmt.Println("a函数: defer函数下面")
	return 0
}

func b() {
	fmt.Println("b函数：defer之后调用")
}

func openFile() {
	f, _ := os.Open("channel.go")
	fmt.Println(f)
	defer f.Close()
}

func main() {
	test := a()
	fmt.Println("main函数：", test)
	//openFile()
}
