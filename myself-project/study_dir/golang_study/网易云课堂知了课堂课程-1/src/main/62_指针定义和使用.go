package main

import "fmt"

func main() {
	// 1. 空指针定义
	var p1 *int         // 空指针
	var p2 *int = nil   // 空指针
	fmt.Println(p1, p2) // <nil> <nil> 两个值都为nil，为空指针

	// 2. 野指针，程序中是不允许存在的
	//var p3 *int
	//*p3 = 56
	//fmt.Println(p3)  // 报错 panic: runtime error: invalid memory address or nil pointer dereference, [signal 0xc0000005 code=0x1 addr=0x0 pc=0x49a7c0]
	// 由于p3没有先指向一个指针，而是直接定义56，指向了一个未知空间，则报错。这就是野指针。

	// 3. 正常操作
	var p4 *int
	a := 56
	p4 = &a
	*p4 = 56
	fmt.Println(p4, *p4)

	// 4. new函数
	//var p5 *int
	//p5 = new(int)
	p5 := new(int)
	*p5 = 57
	fmt.Println(p5, *p5)
}
