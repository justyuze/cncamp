package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)

	defer func() {
		fmt.Println("退出主线程")
	}()
	defer cancel()

	go producer(channel, ctx)

	go consumer(channel, ctx)

	time.Sleep(time.Minute)
}

func producer(channel chan<- int, ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	var n int = 0
	for range ticker.C {
		rand.Seed(time.Now().UnixNano())
		select {
		case <-ctx.Done():
			fmt.Println("主线程发起停止操作，生产者停止生产数据")
			return
		default:
			channel <- n
			fmt.Printf("生产者生产数据 %d, 队列长度为：%d \n", n, len(channel))
		}
		n = n + 1

	}

}

func consumer(channel <-chan int, ctx context.Context) {
	time.Sleep(21 * time.Second)
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("主线程发起停止操作，消费者停止消费数据")
			return
		default:
			n := <-channel
			fmt.Printf("消费者消费数据 %d, 队列长度为：%d \n", n, len(channel))
		}

	}
}
