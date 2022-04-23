package main

import (
	"fmt"
	"strconv"
)

func square(x int) int {
	return x * x
}

func main() {
	fmt.Println("9* 9 :" + strconv.Itoa(square(1)))
}
