package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go rlock()
	go wlock()
	go lock()

	time.Sleep(5 * time.Second)
}

//defer 在退出 rlock函数时才运行
// RLock 不是互斥的，所以可以执行
/*
读锁示例
读锁 -> 0
读锁 -> 1
读锁 -> 2
*/
func rlock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Printf("读锁 -> %d \n", i)
	}
}

// 写锁时互斥的，所以第一次没有释放锁，后面两次都无法再获取写锁
// 写锁 -> 0
func wlock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Printf("写锁 -> %d \n", i)
	}
}

// 标准锁是互斥的
// 锁 -> 0
func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Printf("锁 -> %d \n", i)
	}
}
