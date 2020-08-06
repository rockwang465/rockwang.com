package main

import "fmt"

func main() {
	// 定义string类型的key,数组类型的value
	m := make(map[string][3]int)
	m["小明"] = [3]int{67, 89, 72}
	m["小红"] = [3]int{99, 100, 98}
	m["小华"] = [3]int{33, 29, 46}
	fmt.Println(m) // map[小华:[33 29 46] 小明:[67 89 72] 小红:[99 100 98]]

	for k, v := range m {
		fmt.Println(k, v)
	}

	//fmt.Println(m["张三"])  // 如果没有，则返回默认值[0 0 0 ]
	//同理，如果m的value不是列表，而是int，则返回0； 如果是string，则返回空；如果是布尔，则返回false。

	m2 := make(map[int]string)
	m2[1] = "张三"
	m2[2] = "李四"
	fmt.Println(m2[3]) // 无此key，返回默认值空

	m3 := make(map[int]int)
	m3[1] = 22
	m3[2] = 33
	fmt.Println(m2[3]) // 无此key，返回默认值0

	// 3.可以通过返回值来确认是否存在此key
	value, ok := m3[2]     // ok(可以随便定义一个变量)为获取此key是否存在的返回值,true/false
	fmt.Println(value, ok) // 3 true

	// 4.delete删除map字典
	delete(m3, 1)
	fmt.Println(m3) // map[2:33]
}
