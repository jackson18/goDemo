package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			//  Context 取消的时候，我们就可以得到一个关闭的 chan，关闭的 chan 是可以读取的，所以只要可以读取的时候，就意味着收到 Context 取消的信号
			case <-ctx.Done():
				fmt.Println("任务结束。。。")
				return
			default:
				fmt.Println("任务运行中。。。")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 10)
	fmt.Println("停止任务！！！！")
	cancel()

	time.Sleep(time.Second * 5)
	fmt.Println("程序退出。。。。")

}
