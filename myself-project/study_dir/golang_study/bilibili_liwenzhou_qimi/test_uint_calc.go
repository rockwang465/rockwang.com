package main

import "fmt"

var a uint = 1
var b uint = 2

func main() {
	fmt.Println(a - b) // 18446744073709551615
}

//func main() {
//	b = b - 3
//	fmt.Println("a:", a) // 1
//	fmt.Println("b:", b) // 18446744073709551615
//	fmt.Println(a - b)   // 2
//}
