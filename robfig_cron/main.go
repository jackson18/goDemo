package main

import (
	"github.com/robfig/cron/v3"
	"goDemo/robfig_cron/task"
)

type Cron struct {
	myTask   *task.MyTask
	Schedule *cron.Cron
}

func NewCron(task *task.MyTask) *Cron {
	return &Cron{
		myTask:   task,
		Schedule: cron.New(cron.WithSeconds()), // 让时间支持秒，v3版本默认已经不再是支持到秒级别的定时了
	}
}

func (c *Cron) Start() error {
	// 每天凌晨0点执行一次：0 0 0 * * ?
	// 每隔3秒执行一次：*/3 * * * * *
	// 触发了panic。因为有cron.Recover()保护，后续任务还能执行
	// DelayIfStillRunning触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行
	_, err := c.Schedule.AddJob("*/3 * * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger), cron.Recover(cron.DefaultLogger)).Then(c.myTask))
	if err != nil {
		return err
	}

	c.Schedule.Start()
	return nil
}

func main() {
	s := NewCron(task.NewMyTask("测试任务"))
	s.Start()

	select {}
}
