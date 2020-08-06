package main

import "fmt"

//func main() {
//    // 定义一个空接口
//	var i interface{}
//	fmt.Printf("%T\n", i)  // <nil>
//
//	i = 10
//	fmt.Printf("%T\n", i)  // int
//
//	i = 3.14
//	fmt.Printf("%T\n", i)  // float64
//}

func main() {
	// 2.1 定义一个空接口类型的切片
	var i []interface{}
	i = append(i, "aaa", 10, 3.14)
	fmt.Println(i)
	// 3.1 空接口可以接受任意类型(包括函数名)
	for n := 0; n < len(i); n++ {
		fmt.Println(i[n])
	}
}
