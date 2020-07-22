package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 用于保存随机数的结构体
type job struct {
	value int64
}

// 用于保存随机数 和 随机数的各数之和的结构体
type sum struct {
	job    *job
	sumRes int64
}

var randChan = make(chan *job, 100)
var sumChan = make(chan *sum, 100)
var wg sync.WaitGroup

func randNums(randChan chan<- *job) { // Rock注意1: 这里函数传参的类型定义有点难，要多注意
	defer wg.Done()
	for {
		randVal := rand.Int63() // 生成一个随机数，放入到chan中
		randChan <- &job{value: randVal}
		time.Sleep(time.Millisecond * 500)
	}
}

func sumNums(randChan <-chan *job, sumChan chan<- *sum) {
	//defer wg.Done()
	for {
		job := <-randChan
		jobVal := job.value
		//var sumVal int64
		//sumVal = 0
		sumVal := int64(0)
		for jobVal > 0 {
			tmp := jobVal % 10
			sumVal += tmp
			jobVal = jobVal / 10 // go的普通除法，不存在小数，所以每次都会把余数对应的最后一位值去掉。刚好给下次循环使用。 巧妙的用法。
		}
		sumChan <- &sum{
			job, // 这里要把值赋给job结构体，否则取不到job.value了。
			sumVal,
			//sumRes: sumVal, // 循环完n值后，求得总和，放入进去
		}
	}
}

func main() {
	// 练习题要求:(练习goroutine池)
	// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	// 1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	// 2. 开启24个goroutine从jobChan中取出随机数，拆分为单个数字(拆分单个数字可以通过取余的方式，除10取余，就是最后一位数了，用法很妙)，进行求和，将结果发送到resultChan
	// 3. 主goroutine从resultChan取出结果并打印到终端输出

	wg.Add(1)
	go randNums(randChan)

	// 开启24个goroutine(相当于同时执行24个任务)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go sumNums(randChan, sumChan)
	}

	// 主goroutine从resultChan取出结果并打印到终端输出
	for {
		tmp := <-sumChan
		fmt.Printf("随机数为: %d, 随机数单个数字的和为: %d\n", tmp.job.value, tmp.sumRes)
	}
	wg.Wait()
}
