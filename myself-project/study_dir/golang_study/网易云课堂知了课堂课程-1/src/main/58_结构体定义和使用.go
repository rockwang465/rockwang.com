package main

import "fmt"

type students struct {
	id   int
	name string
	age  int
	sex  string
	addr string
}

func main() {
	// 1. 三种方式为struct结构体赋值:
	// 1.1 顺序初始化(必须按顺序才可以)
	var s1 students = students{1, "rock", 18, "man", "ShangHai"}
	fmt.Println(s1)

	// 1.2 自动推导
	s2 := students{name: "张三", age: 22, sex: "women", id: 2, addr: "ZheJiang"}
	fmt.Println(s2)

	// 1.3 复合类型
	var s3 students
	s3.id = 3
	s3.name = "jacky"
	s3.age = 27
	s3.sex = "man"
	s3.addr = "BeiJing"
	fmt.Println(s3)

	// 2. 结构体的地址和结构体成员的第一个成员的地址相同
	fmt.Printf("%p\n", &s3)    // 0xc00009e080
	fmt.Printf("%p\n", &s3.id) // 0xc00009e080

	// 3. 结构体判断
	fmt.Println(s1 == s2) // 判断两个结构体是否相等

	// 4. 结构体赋值
	var s4 students
	s4 = s3
	fmt.Println(s4)
}
