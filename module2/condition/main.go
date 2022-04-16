package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

// 使用场景， 生产者消费者模式
func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		a := q.Dequeue()
		fmt.Printf("读取到数据 %s \n", a)
	}
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, item)
	fmt.Printf("添加 (%s) 到队列，并通知所有消费者 \n", item)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.queue) == 0 {
		fmt.Println("无数据，等待中。。。。。。")
		q.cond.Wait()
	}

	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
