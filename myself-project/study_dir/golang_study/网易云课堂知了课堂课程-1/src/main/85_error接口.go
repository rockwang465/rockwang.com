package main

import "fmt"

// 定义a b传参，result(int类型) err(error类型)返回值
func test(a, b int) (result int, err error) {
	err = nil   // 先定义为空
	if b == 0 { // 因为b为0，则无法进行除法算法，所以这里加这个判断
		fmt.Println("err=", err)
	} else {
		result = a / b
	}
	return
}

func main() {
	res, err := test(60, 0)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(res)
	}
}
