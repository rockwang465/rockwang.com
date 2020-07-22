package main

import "fmt"

type hero struct {
	name  string
	age   int
	power int
}

func main1() {
	//// 1.1 map字典中value用结构体定义
	//m := make(map[int]hero)
	//m[1] = hero{"钢铁侠",30,5000}
	//m[2] = hero{"奇异博士",33,6000}
	//m[3] = hero{"美国队长",90,4000}
	//fmt.Println(m)  // map[1:{钢铁侠 30 5000} 2:{奇异博士 33 6000} 3:{美国队长 90 4000}]
	//
	//// 1.2 删除
	//delete(m,2)
	//fmt.Println(m)  //map[1:{钢铁侠 30 5000} 3:{美国队长 90 4000}]

	// 2.1 map字典中value用切片结构体定义 -- 难
	m2 := make(map[int][]hero) // map[1:[{钢铁侠 30 5000} {蜘蛛侠 33 2500}]]
	m2[1] = []hero{
		{"钢铁侠", 30, 5000},
		{"蜘蛛侠", 33, 2500},
	}
	fmt.Println(m2) // map[1:[{钢铁侠 30 5000} {蜘蛛侠 33 2500} {黑寡妇 31 4000}]]
	m2[1] = append(m2[1], hero{"黑寡妇", 31, 4000})
	fmt.Println(m2)

	m2[2] = []hero{{"雷神", 1500, 11000}}
	m2[2] = append(m2[2], hero{"惊奇队长", 45, 15000})
	fmt.Println(m2)
}

func main() {
	main1()
}

// m2 := make(map[int][]hero) ，当m2被当函数参数传入时， 为地址传递，会被后面的函数修改值
