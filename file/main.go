package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// 文件信息
	fileInfo, err := os.Stat("file/demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件名：", fileInfo.Name(), ", 文件大小：", fileInfo.Size(), ", 是否目录：", fileInfo.IsDir())

	fileInfo2, err2 := os.Stat("demo2.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("文件名：", fileInfo2.Name(), ", 文件大小：", fileInfo2.Size(), ", 是否目录：", fileInfo2.IsDir())

	fmt.Println("是否绝对路径：", filepath.IsAbs("file/demo.txt"))
	fmt.Println("绝对路径：")
	fmt.Println(filepath.Abs("file/demo.txt"))

	// 创建目录
	err3 := os.MkdirAll("/Users/jacksonqi/Src/go/goDemo/file/test/dir", os.ModePerm)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("多层目录创建成功。。。")

	//// 创建文件（绝对路径）
	//newFile, err4 := os.Create("/Users/jacksonqi/Src/go/goDemo/file/create_file.txt")
	//if err4 != nil {
	//	fmt.Println(err4)
	//	return
	//}
	//fmt.Println("新文件：", newFile)
	//
	//// 创建文件（相对路径）
	//newFile2, err5 := os.Create("file/create_file2.txt")
	//if err5 != nil {
	//	fmt.Println(err5)
	//	return
	//}
	//fmt.Println("新文件：", newFile2)

	// 打开文件
	cfile, err6 := os.OpenFile("file/create_file.txt", os.O_RDWR, os.ModePerm)
	if err6 != nil {
		fmt.Println(err6)
		return
	}
	fmt.Println("打开文件：", cfile)

	// 关闭文件
	cfile.Close()

	// 删除文件或目录
	err7 := os.RemoveAll("/Users/jacksonqi/Src/go/goDemo/file/test/dir")
	if err7 != nil {
		fmt.Println(err7)
		return
	}

}
