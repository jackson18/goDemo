package main

import (
	"fmt"
	"regexp"
)

var (
	// Compile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本。
	validStr, _ = regexp.Compile(`foo.*`)
	// MustCompile类似Compile但会在解析失败时panic，主要用于全局正则表达式变量的安全初始化。
	validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
)

func main() {
	content := "seafood"

	// Match检查b中是否存在匹配pattern的子序列
	fmt.Println(validStr.Match([]byte(content))) // true
	// MatchString类似Match，但匹配对象是字符串
	fmt.Println(validStr.MatchString("seafood"))  // true
	fmt.Println(validID.MatchString("admin[23]")) // true

	// Find返回保管正则表达式在目标数据中的最左侧的一个匹配结果的[]byte切片。如果没有匹配到，会返回nil
	fmt.Printf("%q\n", validStr.Find([]byte(content))) // "food"
	// FindString返回保管正则表达式在目标数据中的最左侧的一个匹配结果的字符串。如果没有匹配到，会返回""；
	// 但如果正则表达式成功匹配了一个空字符串，也会返回""。如果需要区分这种情况，请使用FindStringIndex 或FindStringSubmatch
	fmt.Printf("%q\n", validStr.FindString(content)) // "food"

	// Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片（显然len(loc)==2）。
	// 匹配结果可以通过起止位置对b做切片操作得到：b[loc[0]:loc[1]]。如果没有匹配到，会返回nil
	fmt.Println(validStr.FindStringIndex(content)) // [3 7]
	// Find返回一个保管正则表达式re在b中的最左侧的一个匹配结果以及（可能有的）分组匹配的结果的[]string切片。如果没有匹配到，会返回nil
	fmt.Printf("%q\n", validStr.FindStringSubmatch(content)) // ["food"]

	// ReplaceAllLiteral返回src的一个拷贝，将src中所有re的匹配结果都替换为repl
	// 在替换时，repl中的'$'符号会按照Expand方法的规则进行解释和替换，例如$1会被替换为第一个分组匹配结果
	fmt.Println(validStr.ReplaceAllString(content, "***"))   // sea***
	fmt.Println(validStr.ReplaceAllString(content, "$1###")) // sea###
}
