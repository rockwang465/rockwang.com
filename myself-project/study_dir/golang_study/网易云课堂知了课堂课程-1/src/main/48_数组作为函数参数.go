package main

import "fmt"

//func maopao(arr [10] int) {
func maopao(arr [10]int) [10]int { // return返回值，需要加上[10] int
	// 冒泡排序
	arr_len := len(arr)
	for i := 0; i < arr_len-1; i++ {
		for j := 0; j < arr_len-1; j++ {
			//fmt.Println(j)
			if arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		//fmt.Println(i, arr)
	}
	//fmt.Println(arr)
	return arr
}

func main() {
	var arr [10]int = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//maopao(arr)
	arr = maopao(arr)
	fmt.Println(arr)
}

// 接收参数的函数，需要写明形参的类型，例如：
// test1(a int, b int)
// test2(c [5] int)
