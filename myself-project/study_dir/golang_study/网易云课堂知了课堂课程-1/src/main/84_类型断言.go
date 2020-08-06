package main

import "fmt"

func main8401() {
	var InterF interface{}
	InterF = 10
	//InterF = 10.14

	// 1.1 类型断言的格式:
	// values, ok := element.(type)
	// 值， 值的判断 := 接口变量.(数据类型)
	// ok 值为 true/false

	// 1.2 类型断言定义
	value, ok := InterF.(int)

	// 1.3 if进行类型断言的判断
	if ok { // 如果ok为真
		fmt.Println("数据为整型,值为:", value) // 数据为整型,值为: 10
	} else {
		fmt.Println("数据不为整型,值为:", value) // 数据不为整型,值为: 0
	}
}

func test84() {
	fmt.Println("test84 func")
}

// 2.1 遍历多重判断类型断言
func main8402() {
	// 定义一个切片空接口
	var sliceIF []interface{}
	sliceIF = append(sliceIF, 10, 3.14, "aaa", test84)
	//fmt.Println(sliceIF)

	for _, v := range sliceIF {
		if data, ok := v.(int); ok { // 放在一行判断的
			fmt.Println("数据为整型，值为: ", data)
		} else if data, ok := v.(float64); ok {
			fmt.Println("数据为浮点型，值为: ", data)
		} else if data, ok := v.(func()); ok {
			fmt.Println("数据为函数型，值为: ", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("数据为字符串型，值为: ", data)
		}
	}
}

// 3.1 switch方式进行类型断言的判断
func main() {
	var sliceIF []interface{}
	sliceIF = append(sliceIF, 10, 3.14, "aaa", test84)

	for _, v := range sliceIF {
		switch value := v.(type) {
		case int:
			fmt.Println("数据为整型", value)
		case float64:
			fmt.Println("数据为浮点型", value)
		case string:
			fmt.Println("数据为字符串型", value)
		case func():
			fmt.Println("数据为函数型", value)
			value() // 还能调用函数
		}
	}
}
