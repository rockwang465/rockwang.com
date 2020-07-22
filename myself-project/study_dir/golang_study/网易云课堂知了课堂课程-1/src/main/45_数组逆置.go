package main

import "fmt"

//func main() {
//	var arr [6] int = [6]int{4, 1, 6, 3, 4, 5}
//	// 求上面数组的最大值、最小值、平均值、和
//	max := arr[0]  // 必须给个数组值，否则默认值为0
//	min := arr[0]
//	var avg int
//	var sum int
//
//	for i, _ := range arr {
//		if arr[i] >= max {
//			max = arr[i]  //最大值
//		}
//		if arr[i] <= min{
//			min = arr[i]  // 最小值
//		}
//		sum += arr[i]  // 求和
//	}
//	avg = sum / len(arr)  // 平均值
//	fmt.Printf("最大值: %d ; 最小值: %d ; 平均值: %d ; 和: %d ",max,min,avg,sum)
//}

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var arr [10] int = [10]int{1,4,10,3,6,9,5,7,8,2}
	//将上面数组位置倒过来(逆置)

	//三种for循环介绍
	//1. for 表达式1;表达式2;表达式3{}
	//2. for i,v := range 集合{}
	//3. for { if {} }
	//4. for i < j {}

	//i := 0
	//j := len(arr) - 1
	//方法一:
	//for {
	//	if i < j {
	//		arr[i], arr[j] = arr[j], arr[i]
	//	} else if i == j {
	//		fmt.Printf("i:%d == j:%d \n", i, j)
	//		break
	//	} else if i > j {
	//		fmt.Printf("i:%d > j:%d \n", i, j)
	//		break
	//	}
	//	i ++
	//	j --
	//}
	//fmt.Println(arr)

	//方法二:
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)

}
