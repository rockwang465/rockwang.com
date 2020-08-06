package main

import "fmt"

func main() {
	//var arr [10] int = [10]int{5, 3, 8, 9, 1, 0, 7, 2, 4, 6}
	var arr [10]int = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//var arr [3] int = [3]int{5, 3, 8}

	// 冒泡排序:
	// 从小到大排序:拿第一个和第二个比，比完后小的在前，大的在后
	// 再拿第二个和第三个比， 最后没循环一次，都可以找到除上次外的最大值。
	// 所以当进行 数组长度-1 次，就可以完成冒泡排序了
	arr_len := len(arr)

	for i := 0; i < arr_len-1; i++ {
		for j := 0; j < arr_len-1; j++ {
			//fmt.Println(j)
			if arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(i, arr)
	}
	fmt.Println(arr)
}
