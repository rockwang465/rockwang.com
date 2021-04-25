package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)
	defer close(messages)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(15 * time.Second) // 这里为8-10秒 和 13秒以上有区别,区别是: 10秒以内会打印"child process interrupt...", 超过13秒不打印。
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit")
}

// 问题: 上面sleep部分为8-10秒和13秒以上有区别,区别是:10秒以内会打印"child process interrupt...", 超过13秒不打印。
// 怀疑是gc原因