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

	for i := 1; i <= 10; i++ {
		go producer(channel, ctx, i)
	}

	for i := 1; i <= 5; i++ {
		go consumer(channel, ctx, i)
	}

	time.Sleep(time.Minute)
}

func producer(channel chan<- int, ctx context.Context, no int) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		rand.Seed(time.Now().UnixNano())
		select {
		case <-ctx.Done():
			fmt.Printf("主线程发起停止操作，生产者%d停止生产数据\n", no)
			return
		default:
			n := rand.Intn(100)
			channel <- n
			fmt.Printf("生产者%d生产数据 %d, 队列长度为：%d \n", no, n, len(channel))
		}

	}

}

func consumer(channel <-chan int, ctx context.Context, no int) {
	time.Sleep(time.Second)
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Printf("主线程发起停止操作，消费者%d停止消费数据\n", no)
			return
		default:
			if len(channel) > 0 {
				fmt.Printf("消费者%d开始消费数据\n", no)
				n := <-channel
				fmt.Printf("消费者%d消费数据 %d, 队列长度为：%d \n", no, n, len(channel))
			}

		}

	}
}

/*
输出结果：
主线程发起停止操作，生产者8停止生产数据
主线程发起停止操作，生产者10停止生产数据
主线程发起停止操作，生产者1停止生产数据
主线程发起停止操作，生产者9停止生产数据
主线程发起停止操作，生产者5停止生产数据
主线程发起停止操作，生产者6停止生产数据
主线程发起停止操作，生产者3停止生产数据
主线程发起停止操作，消费者2停止消费数据
主线程发起停止操作，消费者4停止消费数据
主线程发起停止操作，消费者1停止消费数据
主线程发起停止操作，生产者4停止生产数据
主线程发起停止操作，生产者2停止生产数据
主线程发起停止操作，生产者7停止生产数据
退出主线程
*/

/*
消费者3和5没有停止是因为上一次没有读到数据，所以阻塞到那里了，并没有进入下一次循环。所以没有执行到 	case <-ctx.Done():

消费者2开始消费数据
消费者5开始消费数据
消费者3开始消费数据
消费者2消费数据 58, 队列长度为：0
生产者2生产数据 58, 队列长度为：0

以上为生产消费者的最后打印的日志， 可以看到消费者3和5在最后是没有读到数据的
*/
