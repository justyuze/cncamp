package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用场景：主线程启用了多个子线程，且主线程需要等待所有子线程全部完成后才能继续执行。
func main() {
	fmt.Println("开始")
	defer fmt.Println("结束")
	waitByWG()
}

func waitByWG() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			time.Sleep(time.Second)
			fmt.Printf("执行第 %d 个子线程 \n", n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
