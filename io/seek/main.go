package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 断点续传

	srcFile := "demo2.txt"
	destFile := "io/seek/" + srcFile[strings.LastIndex(srcFile, "/")+1:]
	tempFile := destFile + "_bak.txt"

	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file1.Close()
	defer file2.Close()

	// 1.先读取临时文件中的数据，再seek
	file3.Seek(0, io.SeekStart) // 第一个参数表示偏移量，seekStart表示相对于文件开始
	bs := make([]byte, 1024)
	n1, err := file3.Read(bs)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	fmt.Println("临时文件中数据：", count)

	// 2.设置读、写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 2)
	n2 := -1            // 读取的数据量
	n3 := -1            // 写出的数据量
	total := int(count) // 读取的总量

	// 复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件读取完毕，total: ", total)
			file3.Close()
			os.Remove(tempFile)
			break
		}

		n3, err = file2.Write(data[:n2])
		total += n3

		// 将复制的总量，存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		fmt.Println("total: ", total)

		//if total > 5 {
		//	panic("系统故障。。。")
		//}

	}

}
