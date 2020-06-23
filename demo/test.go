package main

import (
	"fmt"
)

type testStu struct {
	Val *int
}

func main() {

	//test := 1 | 2
	//fmt.Println(test)
	slice1 := []int{1, 2, 3}
	//slice1 = testSlice(slice1)
	//fmt.Println(slice1)
	pointer := &slice1
	test := 5

	fmt.Printf("slice：%p pointer: %p \n", slice1, pointer)

	testPointer := &test
	fmt.Printf("test：%p testPointer: %p *testPointer %d\n", test, testPointer, *testPointer)

	var p *int
	x := 5
	p = &x
	fmt.Println("测试*开始")
	fmt.Println(*p)
	fmt.Println(&p)
	fmt.Println(*(&p))
	fmt.Println(*(&(*p)))

	fmt.Println("测试*结束")

	fmt.Println("测试...开始")
	result := make([]int, 0)
	result = append(result, 0)
	fmt.Println(result)
	left := []int{1, 2, 3}
	result = append(result, left...)
	fmt.Println(result)
	fmt.Println("测试...结束")

	x = 10
	fmt.Println(*p)

	testPointInt(p)
	fmt.Println(*p)

	st := testStu{}
	*st.Val = 10
}
func testPointInt(val *int) {
	*val += 5
}
func testSlice(arr []int) []int {
	arr[1] = 2
	arr = append(arr, 0)
	return arr
}
