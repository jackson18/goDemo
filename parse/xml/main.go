package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Address string   `xml:"address,attr"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Social  Social   `xml:"social"`
}
type Social struct {
	XMLName xml.Name `xml:"social"`
	Mobile  string   `xml:"mobile"`
	Email   string   `xml:"email"`
}

func readXml() {
	xmlFile, err := os.Open("parse/xml/user.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("successfully opened users.xml")
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Address: ", users.Users[i].Address)
		fmt.Println("User Name: ", users.Users[i].Name)
		fmt.Println("User Age: ", users.Users[i].Age)
		fmt.Println("Facebook Url: ", users.Users[i].Social.Email)
	}
}

func writeXml() {
	dongdong := User{Address: "chengdu", Name: "冬哥", Age: 30, Social: Social{Email: "dongdong@163.com", Mobile: "1155555211111"}}
	zhengge := User{Address: "beijing", Name: "郑哥", Age: 24, Social: Social{Email: "zhengge@163.com", Mobile: "1112224566211111"}}
	v := &Users{Users: []User{zhengge, dongdong}}

	result, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(result))
	fileName := "parse/xml/newUser.xml"
	err = ioutil.WriteFile(fileName, result, 0777)
	if err != nil {
		panic("Unable to write data into the file")
	}
}

func main() {
	readXml()
	writeXml()
}
