package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	read()
	read2()
	read3()
	//read4()
	read5()
	write()
}

func write() {
	fileName := "bufio/demo.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString("hello go")
	w.Flush()
}

func read5() {
	b := bufio.NewReader(os.Stdin)
	s, _ := b.ReadString('\n')
	fmt.Println("读取到的键盘输入(以换行为分隔符)：", s)
}

func read4() {
	s := ""
	fmt.Scanln(&s)
	fmt.Println("读取到的键盘输入(有空格就会中断，不好用)：", s)
}

func read3() {
	fileName := "demo2.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b1 := bufio.NewReader(file)
	for {
		s1, err := b1.ReadString('\n')
		fmt.Println("ReadString: ", s1)
		if err == io.EOF {
			fmt.Println("读取完毕。。。")
			break
		}
	}
	fmt.Println()
}

// 不建议使用
func read2() {
	fileName := "demo2.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b1 := bufio.NewReader(file)
	data, flag, err := b1.ReadLine()
	fmt.Println("ReadLine: ", flag, string(data))
	fmt.Println()
}

func read() {
	fileName := "demo2.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b1 := bufio.NewReader(file)
	bs := make([]byte, 1024)
	count, err := b1.Read(bs)
	fmt.Println("Reader: ", count, string(bs[:count]))
	fmt.Println()
}
