package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {

	p := "D:/123/456"
	d := "D:\\123\\456"
	// 获取最后的路径段的值
	fmt.Println("path.Base: ", path.Base(p))
	// 获取最后的路径段的值，如果是windows风格的路径，需要转化一下
	fmt.Println("path.Base: ", path.Base(d), path.Base(strings.ReplaceAll(d, "\\", "/")))

	// 是否是一个绝对路径
	fmt.Println("path.IsABS: ", path.IsAbs("/root/users"), path.IsAbs("./data/user"))

	// 将多个路径合并成一个单一路径,会自动添加斜杠
	fmt.Println("path.Join: ", path.Join("/root", "users", "data"))

	// 将路径从最后一个斜杠位置分成两部分
	dir, action := path.Split(p)
	fmt.Println("path.Split: ", dir, action)

	// 返回路径除去最后一个路径元素的部分
	// 如果路径是空字符串，会返回 .
	// 如果路径由 1 到多个斜杠后跟 0 到多个非斜杠字符组成，会返回 /
	fmt.Println("path.Dir: ", path.Dir(p), path.Dir(""), path.Dir("/123"), path.Dir("///123"))

	// 返回path文件扩展名
	fmt.Println("path.Ext: ", path.Ext("d:/go/aa.png"))

}
