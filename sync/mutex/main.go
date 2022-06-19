package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var tickets = 10
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(2)

	go saleTicket("张三")
	go saleTicket("李四")

	wg.Wait()
	fmt.Println("over...")

}

func saleTicket(name string) {
	for {
		mutex.Lock()

		if tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, ": ", tickets)
			tickets--
		} else {
			fmt.Println(name, ", 结束卖票。。。")
			mutex.Unlock()
			break
		}

		mutex.Unlock()
	}
	wg.Done()
}
