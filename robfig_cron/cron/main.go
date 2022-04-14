package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	//每天10点执行一次：0 0 10 * * *
	//每隔10分钟执行一次：0 */10 * * *
	//每月1号凌晨3点执行一次：0 0 3 1 * ?
	//每月最后一天23点30分执行一次：0 30 23 L * ?
	//每周周六凌晨3点实行一次：0 0 3 ? * L
	//在30分、50分执行一次：0 30，50 * * * ?

	spec := "*/3 * * * * *"
	c := cron.New(cron.WithSeconds())
	c.AddFunc(spec, func() {
		fmt.Println("execute")
	})
	go c.Start()
	defer c.Stop()
	select {}
}
