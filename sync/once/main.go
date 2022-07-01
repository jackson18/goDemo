package main

import (
	"fmt"
	"sync"
)

func main() {

	once := sync.Once{}
	count := 0
	for i := 0; i <= 10; i++ {
		once.Do(func() {
			fmt.Println("count++")
			count++
		})
		fmt.Println("i: ", i, ", count: ", count)
	}

	fmt.Println(count)

}
