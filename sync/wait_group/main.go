package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printOne(&wg)
	go printTwo(&wg)

	wg.Wait()
	fmt.Print("over...")
}

func printOne(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}

	wg.Done()
}

func printTwo(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}

	wg.Done()
}
