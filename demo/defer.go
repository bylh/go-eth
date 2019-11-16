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
	//test := a()
	//fmt.Println("main函数：", test)
	////openFile()

	for i := 0; i < 5; i++ {
		fmt.Println(i, "before")
		defer fmt.Println("defer runs", i)
		fmt.Println(i, "after")
	}

	{
		defer fmt.Println("defer runs 5")
		defer fmt.Println("defer runs 6")

		fmt.Println("block ends")
	}

	fmt.Println("main ends")

	// defer 执行顺序 defer runs 6 5 4 3 2 1 可以看出，def的执行都在最后，且def执行排序，在代码中越靠后，越先执行

	// defer在函数return后执行 一般用来释放资源，否则每个返回的地方都要手动调用释放资源的代码
}
