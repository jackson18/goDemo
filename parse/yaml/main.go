package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Users struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Address string   `yaml:"address"`
	Hobby   []string `yaml:"hobby"`
}

func readYaml() {
	file, err := ioutil.ReadFile("parse/yaml/test.yaml")
	if err != nil {
		fmt.Println("read file error...", err)
		return
	}

	var data = make([]Users, 10)
	err2 := yaml.Unmarshal(file, &data)
	if err2 != nil {
		fmt.Println("yaml Unmarshal error...", err2)
		return
	}

	for _, v := range data {
		fmt.Println(v)
	}
}

func writeYaml() {
	huazai := Users{
		Name:    "华子",
		Age:     28,
		Address: "shenzhen",
		Hobby:   []string{"王者荣耀"},
	}
	qiaoke := Users{
		Name:    "乔克",
		Age:     30,
		Address: "chongqing",
		Hobby:   []string{"阅读", "王者荣耀"},
	}

	userlist := []Users{huazai, qiaoke}

	yamlData, err := yaml.Marshal(&userlist)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(string(yamlData))
	fileName := "parse/yaml/new.yaml"
	err = ioutil.WriteFile(fileName, yamlData, 0777)
	if err != nil {
		panic("Unable to write data into the file")
	}
}

func main() {
	readYaml()
	writeYaml()
}
