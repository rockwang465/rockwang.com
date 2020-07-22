package main

import "fmt"

//func main() {
//	// 数组定义
//	var a [10] int
//	fmt.Println(len(a))
//	fmt.Println(a)  // 结果: [0 0 0 0 0 0 0 0 0 0]
//
//	// 数组赋值
//	for i := 0; i<len(a); i++ {
//		a[i] = i*i
//	}
//	fmt.Println(a)  // 结果: [0 1 2 3 4 5 6 7 8 9]
//
//	// 取数组的值
//	//for i :=0; i<len(a); i++ {
//	//	fmt.Println(a[i])
//	//}
//
//	// range来循环a数组，_为索引，匿名变量去除；v为值
//	for _,v := range a{
//		fmt.Println(v)
//	}
//
//	// 其他类型的定义
//	var b [11] string
//	fmt.Println(b)  // 字符串类型默认为空，结果为: [          ]
//
//	var c [12] float64
//	fmt.Println(c)  // 浮点类型默认为0，结果为[0 0 0 0 0 0 0 0 0 0 0 0]
//
//	var d [5] bool
//	fmt.Println(d)  // 布尔类型默认为false，结果为[false false false false false]
//}

//func main() {
//	// 数组初始化
//	var a [5] int = [5]int{1, 2, 3, 4, 5}
//	fmt.Println(a)
//
//	// 自动推导
//	b := [5]int{1, 2, 3, 4, 5}
//	fmt.Println(b)
//
//	// 部分初始化
//	c := [5] int{1, 2, 3}
//	fmt.Println(c)
//
//	// 指定某个元素初始化
//	d := [5]int{2: 10, 4: 20} // 指定索引为2的值为10， 索引为4的值为20
//	fmt.Println(d)
//
//	// 定义长度不定的数组
//	f := [...]int{1, 2, 3}
//	fmt.Println(len(f), f)
//}

func main() {
	// 变量赋值数组
	arr := [5]int{1, 2, 3, 4, 5}

	arr2 := arr
	fmt.Println(arr)
	fmt.Println(arr2)

	// 数组的格式与数组的内存地址
	fmt.Printf("%T\n", arr)     // 打印arr的格式，结果为 [5]int
	fmt.Printf("%p\n", &arr)    // 获取arr的内存地址，结果为0xc00000c2d0
	fmt.Printf("%p\n", &arr[0]) // 获取元素的内存地址，结果为0xc00000c2d0，发现数组名和数组的第一个元素的地址是一样的
	fmt.Printf("%p\n", &arr[1]) // 获取元素的内存地址，结果为0xc00000c2d8，和其他的元素地址不同
	fmt.Printf("%p\n", &arr[2]) // 获取元素的内存地址，结果为0xc00000c2e0，和其他的元素地址不同
	fmt.Printf("%p\n", &arr[3])
	fmt.Printf("%p\n", &arr[4])
}
