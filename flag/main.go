package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	//第一个参数：存放值的参数地址
	//第二个参数：命令行参数的名称
	//第三个参数：命令行不输入时的默认值
	//第四个参数：该参数的描述信息，help命令时会显示

	// 接收命令参数
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

// 调用：go run main.go -name="jacksonqi" -age=30
// 查看源码文件的参数说明：go run main.go --help
func main() {
	//参数	功能
	//name	命令行参数名称，比如 -b, -help
	//value	默认值，未显式指定的参数，给出隐式的默认值，比如本例中-b未给出的话，*b=false
	//usage	提示信息，如果给出的参数不正确或者需要查看帮助 -help，那么会给出这里指定的字符串
	age := flag.Int64("age", 20, "param age")

	// 解析命令参数，并把他们的值赋值给相应的变量
	flag.Parse()

	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("age: ", *age)
}
