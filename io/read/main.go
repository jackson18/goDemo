package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 读取文件数据
	file, err := os.Open("demo2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	bs := make([]byte, 1024)
	n := -1
	for true {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("文件读取完成。。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
