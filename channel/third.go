package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square2(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
	fmt.Println("never Print")
}

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("square:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
