package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writerHello(writer Writer) {
	writer("hello world")
}

func main() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println("Fail to create file")
		return
	}
	defer f.Close()

	writerHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
