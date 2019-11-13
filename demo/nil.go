package main

import "fmt"

type FetchData struct {
	Data []map[string]interface{}
}

func main() {
	// var t []string  // 声明空数组方式一，此时t为nil
	// t := []string{} // 声明空数组方式二， 此时t不为nil 等价于var a0 []int = make([]int, 0)

	var t []string
	println("测试t", len(t))
	fetchData := FetchData{
		Data: []map[string]interface{}{}, // 声明空数组
	}

	fmt.Println(len(fetchData.Data))
}
