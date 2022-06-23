package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 获取当前时间
	fmt.Println("获取当前时间: ", time.Now())

	// 指定时间
	fmt.Println("指定时间: ", time.Date(2022, 4, 17, 10, 12, 30, 0, time.Local))

	// 创建本地时间，并携带时区
	fmt.Println("创建本地时间，并携带时区：", time.Unix(time.Now().Unix(), 0).String())

	// 时间转str
	t1 := time.Now()
	fmt.Println("时间转str: ", t1.Format("2006-1-2 15:04:05"))
	fmt.Println("时间转str :", t1.Format("2006/01/02"))

	// str转时间
	s := "2022年04月17日"
	t2, _ := time.Parse("2006年01月02日", s)
	fmt.Println("str转时间：", t2.Year(), t2.Month(), t2.Day())

	// 根据当前时间，获取指定内容
	year, month, day := time.Now().Date()
	fmt.Println("year: ", year, ", month: ", month, ", day: ", day)

	hour, min, sec := time.Now().Clock()
	fmt.Println("hour: ", hour, ", min: ", min, ", sec: ", sec)

	// 时间戳
	fmt.Println("时间戳(单位秒)：", time.Now().Unix())
	fmt.Println("时间戳(单位纳秒)：", time.Now().UnixNano())

	// 时间间隔
	fmt.Println("时间间隔加一分钟：", time.Now().Add(time.Minute))

	// 随机数
	// 改变随机数生成器的种子
	rand.Seed(time.Now().UnixNano())
	fmt.Println("随机数（1~10）：", rand.Intn(10)+1)

}
