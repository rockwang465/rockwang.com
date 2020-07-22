package main

import "fmt"

func main() {
	// 1. map字典定义: map[键类型]值类型
	// 1.1 var定义
	//var m1 map[int]string
	// 1.2 自动推导
	//m2 := make(map[int]string)
	// 注意: 虽然不能使用cap容量函数，但可以定义容量值，这里定义了3的容量。
	//m3 := make(map[int]string, 3)
	// 1.3 初始化
	//m4 := map[int]string{1:"张三",2:"李四"}  // map[1:张三 2:李四]

	//fmt.Println(m)  // map[]
	// 注意: 在字典类型中，只能用len长度函数，不能有cap容量函数。

	// 2. 赋值
	m6 := make(map[int]string)
	m6[1] = "张三"
	fmt.Println(len(m6)) // 结果:1, 长度自动变化，所以不用定义长度
	m6[2] = "李四"
	fmt.Println(len(m6)) // 结果:2
	m6[3] = "王五"
	m6[4] = "陈六"

	fmt.Println(m6)    // map[1:张三 2:李四 4:陈六 3:王五] 字典是无序的
	fmt.Println(m6[3]) // 3.取值， 结果为:王五
}
