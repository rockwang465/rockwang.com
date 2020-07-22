package main

import "fmt"

func main() {
	var arr [2][3]int

	fmt.Println(arr)         // [[0 0 0] [0 0 0]]
	fmt.Println(len(arr))    // 查行数， 结果:2
	fmt.Println(len(arr[0])) //查第0行的列数，结果:3
	fmt.Println(len(arr[1])) //查第1行的列数，结果:3，这里查列数都一样

	arr[0][1] = 123  // 数组的第0行的索引为1的值改为123，结果:[[0 123 0] [0 0 0]]
	arr[1][2] = 456  // 数组的第1行的索引为2的值改为456，结果:[[0 123 0] [0 0 456]]
	fmt.Println(arr) // 结果为: [[0 123 0] [0 0 456]]
}

func main() {
	// 1.全部初始化
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(b)
	// 2.部分初始化
	c := [3][4]int{{5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(c)
}
