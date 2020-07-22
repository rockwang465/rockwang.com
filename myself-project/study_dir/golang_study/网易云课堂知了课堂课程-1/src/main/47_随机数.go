package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.导入头文件(导入模块) math/rand
	//2.随机数种子(实例化)
	rand.Seed(1)                     // (很少用)1是个固定的数字，所以随机的值会一直固定不变
	rand.Seed(time.Now().UnixNano()) // (常用)时间是一直变化的，所以随机的值也会一直变化
	//3.创建随机数(生成随机数)
	fmt.Println(rand.Int())      // 5577006791947779410
	fmt.Println(rand.Intn(1000)) //生成10以内的数字(0-9)

	// 双色球案例
	// 红球 1-33编号，选择6个，不能重复
	// 蓝球 1-16编号，选择1个，可以和红球重复

	var red_ball [6]int
	var blue_ball [1]int
	var rand_num int
	rand.Seed(time.Now().UnixNano()) // 注意，千万别把这个放到循环内了，否则值都一样，不会变化

	for i := 0; i < len(red_ball); i++ {
		for j := 0; j < len(red_ball); j++ {
			rand_num = rand.Intn(33) + 1 //+1，解决生成0-32之间的数中有0和没有33的问题。
			if rand_num == red_ball[j] { // 如果当前随机数和之前的相同
				rand_num = rand.Intn(33) + 1 // 重新生成
			}
		}
		red_ball[i] = rand_num
		//time.Sleep(1 * time.Second)
	}
	//time.Sleep(1 * time.Second)
	rand_num = rand.Intn(13)
	blue_ball[0] = rand_num + 1
	fmt.Printf("红球: %d ; 篮球: %d", red_ball, blue_ball)
}
