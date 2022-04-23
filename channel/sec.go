package main

import (
	"fmt"
	"time"
)

func main() {

	go square()
	ch := make(chan int, 2)
	ch <- 9
	fmt.Println("never Print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
