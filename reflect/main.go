package main

import (
	"fmt"
	"reflect"
)

type myInt int

const (
	Zero myInt = 0
)

type Student struct {
	Name string
	Age int `json:"age"`
}

func (s Student) GetName() string {
	return s.Name
}

func main() {

	c := make([]int, 10)
	c[0] = 100
	fmt.Println("type: ", reflect.TypeOf(c), reflect.TypeOf(c).Kind(), reflect.ValueOf(c), reflect.ValueOf(c).Kind())

	var stu Student
	// 获取反射类型对象
	typeOfStu := reflect.TypeOf(stu)
	fmt.Println("stu：", typeOfStu.Name(), typeOfStu.Kind())
	fmt.Println("zeor: ", reflect.TypeOf(Zero).Name(), reflect.TypeOf(Zero).Kind())

	s := new(Student)
	s.Name = "张三"
	s.Age = 18
	// Go语言的反射中对所有指针变量的种类都是 Ptr
	fmt.Println("s.Name: ", reflect.TypeOf(s).Name(), "s.Kind: ", reflect.TypeOf(s).Kind())
	// reflect.Elem() - 通过反射获取指针指向的元素类型
	fmt.Println("s.Name: ", reflect.TypeOf(s).Elem().Name(), "s.Kind: ", reflect.TypeOf(s).Elem().Kind())

	ss := Student{
		Name: "张三",
		Age: 20,
	}
	// 获取反射类型对象
	typeOfSS := reflect.TypeOf(ss)
	// 遍历所有成员
	for i := 0; i < typeOfSS.NumField(); i++ {
		f := typeOfSS.Field(i)
		fmt.Println("fieldName: ", f.Name)
	}
	// 通过字段名找字段类型信息
	if ageType, ok := typeOfSS.FieldByName("Age"); ok {
		fmt.Println("ageType: ", ageType.Tag.Get("json"))
	}

	// 获取字段值
	v := reflect.ValueOf(ss)
	fmt.Println("age.value: ", v.FieldByName("Age"))
	fmt.Println("name.value: ", v.FieldByName("Name"))

	// 判断反射值的空和有效性
	var num *int
	fmt.Println("IsNil: ", reflect.ValueOf(num).IsNil())
	// 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等
	fmt.Println("IsValid: ", reflect.ValueOf(nil).IsValid())
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的字段 IsValid: ", reflect.ValueOf(ss).MethodByName("").IsValid())

	// 修改字段值
	val := reflect.ValueOf(&ss)
	name := val.Elem().FieldByName("Name")
	name.SetString("李四")
	fmt.Println("修改后的name: ", name.String(), ss.Name)

	// 通过类型创建实例
	var a int
	typeOfA := reflect.TypeOf(a)
	aIns := reflect.New(typeOfA)
	fmt.Println("aIns: ", aIns.Type(), aIns.Kind())

	// 反射调用方法
	s1 := Student{Name: "王五"}
	rVal := reflect.ValueOf(s1)
	fmt.Println(rVal.NumMethod())
	//构造函数参数，传入两个整形值
	paramList := []reflect.Value{}
	// 调用结构体的第一个方法Method(0)
	// 注意:在反射值对象中方法索引的顺序并不是结构体方法定义的先后顺序
	// 而是根据方法的ASCII码值来从小到大排序，所以Dec排在第一个，也就是Method(0)
	result := rVal.Method(0).Call(paramList)
	fmt.Println("result: ", result)

}
