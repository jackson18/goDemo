package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	read()
	write()
	readAll()
	readDir()
	tempDirAndFile()
	recursion("io", 0)
}

// 递归遍历目录
func recursion(dir string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}

	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range fs {
		name := dir + "/" + f.Name()
		fmt.Printf("%s%s\n", s, name)
		if f.IsDir() {
			recursion(name, level+1)
		}
	}
}

func tempDirAndFile() {
	// 创建临时目录
	dir, err := ioutil.TempDir("ioutil", "temp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	fmt.Println(dir)

	// 创建临时文件
	file, err := ioutil.TempFile(dir, "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())
}

// readDir只能查下面一级的所有文件及目录
func readDir() {
	dirName := "io"
	fs, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range fs {
		fmt.Println(fs[i].Name(), fs[i].IsDir(), fs[i].Mode())
	}
}

func readAll() {
	s := "好好学习，天天向上"
	r := strings.NewReader(s)
	data, _ := ioutil.ReadAll(r)
	fmt.Println("ReadAll: ", string(data))
}

func read() {
	fileName := "demo2.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ReadFile: ", string(data))
}

func write() {
	fileName := "ioutil/a.txt"
	s := "统一大市场"
	ioutil.WriteFile(fileName, []byte(s), os.ModePerm)
}
