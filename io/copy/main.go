package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	srcFile  = "demo2.txt"
	destFile = "io/copy/copy.txt"
)

func main() {
	//_, err := copyFile1(srcFile, destFile)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//_, err := copyFile2(srcFile, destFile)
	//if err != nil {
	//	fmt.Println(err)
	//}

	_, err := copyFile3(srcFile, destFile)
	if err != nil {
		fmt.Println(err)
	}
}

// 如果文件过大，容易内存溢出， 不适用于大文件
func copyFile3(src, dest string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}

	return len(bs), nil
}

func copyFile2(src, dest string) (int64, error) {
	file1, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}

func copyFile1(src, dest string) (int, error) {
	file1, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完成。。。")
			break
		} else if err != nil {
			fmt.Println("拷贝异常。。。")
			return total, err
		}

		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
