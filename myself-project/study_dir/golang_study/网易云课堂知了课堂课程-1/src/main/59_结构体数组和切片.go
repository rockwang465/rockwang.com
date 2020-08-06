package main

import "fmt"

type Student struct {
	id    int
	name  string
	score int
}

// 例1.结构体数组的操作
// 2.排序操作
func Sort(arr [3]Student) { // 注意: 传入形参类型时，要加数组的数量3
	fmt.Println("函数执行前打印: ", arr) // [{1 李白 100} {2 杜甫 93} {3 王维 97}]
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].score < arr[j+1].score {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("函数执行后打印: ", arr) // [{1 李白 100} {3 王维 97} {2 杜甫 93}]
}

// 1.基础操作
func main() {
	var arr [3]Student = [3]Student{
		{1, "李白", 100},
		{2, "杜甫", 100},
		{3, "王维", 100},
	}
	//fmt.Println(arr)  // 结构体数组: [{1 李白 100} {2 杜甫 100} {3 王维 100}]

	//for i:=0; i<len(arr); i++{
	//	fmt.Println(arr[i])  // {1 李白 100}  {2 杜甫 100}  {3 王维 100}
	//}

	// 修改结构体数组成员信息
	arr[1].score = 93
	arr[2].score = 97
	//fmt.Println(arr)  // [{1 李白 100} {2 杜甫 98} {3 王维 97}]

	// 直接append是无法操作的，因为数组的长度是固定的。 只有用切片才可以append
	//arr = append(arr,Student{4,"李清照", 101})

	// 排序:按照score分数排序
	Sort(arr)
	fmt.Println("main主函数打印: ", arr) // 不是地址传入，所以结果不会被Sort排序函数修改。
}

//// 例2. 结构体切片操作
//// 2.排序操作
//func Sort(arr []Student) { // 注意: 传入形参类型时，要加数组的数量3
//	fmt.Println("函数执行前打印: ",arr)  // [{1 李白 100} {2 杜甫 93} {3 王维 97}]
//	for i := 0; i < len(arr)-1; i++ {
//		for j := 0; j < len(arr)-1; j++ {
//			if arr[j].score < arr[j+1].score {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//	fmt.Println("函数执行后打印: ",arr)  // [{1 李白 100} {3 王维 97} {2 杜甫 93}]
//}
//
//
//// 1.基础操作
//func main() {
//	var arr []Student = []Student{
//		{1, "李白", 100},
//		{2, "杜甫", 100},
//		{3, "王维", 100},
//	}
//
//	// 修改结构体数组成员信息
//	arr[1].score = 93
//	arr[2].score = 97
//	//fmt.Println(arr)  // [{1 李白 100} {2 杜甫 98} {3 王维 97}]
//
//	// 直接append是无法操作的，因为数组的长度是固定的。 只有用切片才可以append
//	arr = append(arr,Student{4,"李清照", 101})
//
//	// 排序:按照score分数排序
//	Sort(arr)
//	fmt.Println("main主函数打印: ",arr)  // arr已经被Sort排序函数修改，因为结构体数组是地址传值
//}
