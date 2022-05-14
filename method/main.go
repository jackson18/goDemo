package main

import (
	"fmt"
	"reflect"
)

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "hi " + a.name
	return a.name
}

func main() {

	a := A{name: "张三"}
	// 语法糖
	fmt.Println(a.Name())
	// a是方法接收者
	fmt.Println(A.Name(a))

	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).IsValid())

}
