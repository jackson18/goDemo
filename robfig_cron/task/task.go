package task

import (
	"fmt"
	"math/rand"
	"time"
)

type MyTask struct {
	name string
}

func NewMyTask(name string) *MyTask {
	return &MyTask{name: name}
}

// 实现Job接口中的Run方法后，可调用cron对象的AddJob()方法将MyTask对象添加到定时管理器中
func (t *MyTask) Run() {
	fmt.Println("myTask is running...")
	num := rand.Intn(3)
	if num == 1 {
		panic("程序异常。。。。")
	} else if num == 2 {
		fmt.Println("延迟5秒。。。。")
		time.Sleep(5 * time.Second)
	}
}
