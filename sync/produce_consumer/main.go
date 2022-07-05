package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发消费示例
func main() {

	wg := sync.WaitGroup{}

	taskCount := 100
	// 10个消费者
	curConsumerNum := 10

	taskChan := make(chan int, curConsumerNum)
	defer close(taskChan)

	for i := 0; i < curConsumerNum; i++ {
		go func() {
			for task := range taskChan {
				doTask(&wg, task)
			}
		}()
	}

	for i := 1; i <= taskCount; i++ {
		wg.Add(1)
		taskChan <- i
	}

	wg.Wait()

	fmt.Println("main over!!!")
}

func doTask(wg *sync.WaitGroup, index int) {
	fmt.Println(index, " : doTask start...")
	time.Sleep(time.Second * 1)
	fmt.Println(index, " : doTask done")
	wg.Done()
}
