package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	copy(s2, s1)      // 将s1内容赋值给s2
	copy(s2, s1[0:3]) // 将s1中的切片内容赋值给s2
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%p\n", s1) // 地址互相不同
	fmt.Printf("%p\n", s2) // 地址互相不同

	// 冒泡排序
	s3 := []int{9, 7, 5, 3, 1, 8, 6, 4, 2, 0}
	for i := 0; i < len(s3)-1; i++ {
		for j := 0; j < len(s3)-1; j++ {
			if s3[j] > s3[j+1] {
				s3[j], s3[j+1] = s3[j+1], s3[j]
			}
		}
	}
	fmt.Println(s3)

}
