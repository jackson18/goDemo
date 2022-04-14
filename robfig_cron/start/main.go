package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// 每秒触发一次
	/**
	@yearly：也可以写作@annually，表示每年第一天的 0 点。等价于0 0 1 1 *；
	@monthly：表示每月第一天的 0 点。等价于0 0 1 * *；
	@weekly：表示每周第一天的 0 点，注意第一天为周日，即周六结束，周日开始的那个 0 点。等价于0 0 * * 0；
	@daily：也可以写作@midnight，表示每天 0 点。等价于0 0 * * *；
	@hourly：表示每小时的开始。等价于0 * * * *。
	*/
	c.AddFunc("@every 1s", func() {
		fmt.Println("task start in 1 seconds")
	})

	c.Start()

	select {}
}
