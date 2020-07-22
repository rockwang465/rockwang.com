package main

import "fmt"

func main8601() {
	fmt.Println("hello1")
	fmt.Println("hello2")
	panic("helloooooooo")
	fmt.Println("hello3")

	// 结果:
	//	hello1
	//	hello2
	//  panic: helloooooooo
}

func testSlice(i int) {
	var arr [3]int
	arr[i] = 999
	fmt.Println(arr)
}

func main() {
	testSlice(3)
	//报错: panic: runtime error: index out of range [3] with length 3
}
