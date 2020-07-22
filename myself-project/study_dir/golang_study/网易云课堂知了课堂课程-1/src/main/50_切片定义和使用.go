package main

import "fmt"

//func main() {
//	// 1.定义切片
//	var slice [] int // 定义了一个空切片，slice命名随意。
//
//	slice = append(slice, 1, 2, 3, 4, 5) // 追加一个1，并必须赋值
//	fmt.Println(slice)
//	fmt.Println(len(slice)) //获取切片的长度
//	fmt.Println(cap(slice)) //获取切片的容量
//
//	fmt.Printf("%p\n", slice)  // 切片地址 0xc000094030
//	fmt.Printf("%p\n", &slice[0])  // 切片第0个索引地址0xc000094030,和切片地址相同
//	fmt.Printf("%p\n", &slice[1])  // 切片第1个索引地址0xc000094038,和其他不同
//
//	// 2.自动推导定义切片
//	var s1 [] int = [] int{1,2,3}
//	s2 := [] int{1,2,3,4,5,6}
//	fmt.Println(s1)
//	fmt.Println(s2)
//}

func main() {
	s := make([]int, 0, 1) //创建一个切片长度为0，容量为1
	oldcap := cap(s)
	for i := 0; i < 2000; i++ {
		s = append(s, i)
		newcap := cap(s)
		if oldcap < newcap {
			fmt.Printf("oldcap: %d  ------ newcap: %d\n", oldcap, newcap)
			oldcap = newcap
		}
	}
}
