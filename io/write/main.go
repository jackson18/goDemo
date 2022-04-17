package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("demo2.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	n, err2 := file.Write([]byte("hello jacksonqi"))
	file.WriteString("\ngood morning!")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(n)

}
