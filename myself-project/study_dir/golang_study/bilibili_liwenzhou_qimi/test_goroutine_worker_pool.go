package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	Job int
	Sum int
}

var (
	//wg  sync.WaitGroup
	jobs    = make(chan int, 20)
	results = make(chan Result, 20)
)

func calc(v int) int {
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(5)
	y := rand.Intn(10)
	return (v + x) * y
}

func do(j int, wg *sync.WaitGroup) {
	for v := range jobs {
		//results <- Result{Job: v, Sum: v * 2} // 这里还可以改一下，加个func
		results <- Result{Job: v, Sum: calc(v)}
		fmt.Printf("%v : do something ...\n", j)
		time.Sleep(time.Second * 2)
		wg.Done()
	}
}

func importJob(jobSize int) {
	for i := 0; i < jobSize; i++ {
		jobs <- i
	}
	close(jobs)
}

func startWorkerPool(workerSize int) {
	var wg sync.WaitGroup
	for j := 0; j < workerSize; j++ {
		wg.Add(1)
		go do(j, &wg)
	}
	wg.Wait()
	close(results)
}

func printResult(done chan<- bool) {
	for {
		v, ok := <-results
		if !ok {
			done <- true
			break
		}
		fmt.Printf("print result : %v\n", v)
	}
}

func main() {
	jobSize := 10 // 任务数量
	go importJob(jobSize)

	// 开启5个worker pool
	workerSize := 5 // 工作池
	go startWorkerPool(workerSize)

	// 打印结果
	done := make(chan bool)
	go printResult(done)
	<-done // 使用的亮点，卡住进程。只有worker poll都结束，close了results，才能给done一个值。
}
