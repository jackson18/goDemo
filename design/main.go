package main

import (
	"fmt"
)

type dealData func(req string) (bool, error)

func dealWeixinData(req string) (bool, error) {
	fmt.Println("dealWeixinData: ", req)
	return true, nil
}

func dealAliData(req string) (bool, error) {
	fmt.Println("dealAliData: ", req)
	return true, nil
}

var (
	mm = map[string]dealData{
		"weixin": dealWeixinData,
		"ali":    dealAliData,
	}
)

func main() {
	deal, ok := mm["weixin"]
	if !ok {
		fmt.Println("类型不合法。。。")
		return
	}

	res, _ := deal("request")
	fmt.Println(res)
}
