package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// s[起点位置:结束位置:计算容量的最大值]
	// s2切片长度 = 结束位置 - 起点位置
	// s2切片容量 = 计算容量的最大值 - 起点位置
	s2 := s[0:2:6]
	fmt.Println(len(s2)) // 长度: 2 - 0 = 2
	fmt.Println(cap(s2)) // 容量: 6 - 0 = 6

	s2[1] = 999            // 修改新的切片
	fmt.Println(s2)        // 结果: [0 999]
	fmt.Println(s)         // 结果: [0 999 2 3 4 5 6 7]， 这里发现被改动了
	fmt.Printf("%p\n", s)  // 0xc000084080
	fmt.Printf("%p\n", s2) // 0xc000084080
	// 所以，切片后的值修改，会影响之前的切片。
	// 因为实际还是操作之前的切片
}
