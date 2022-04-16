package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000)
	cnt := 1
	fmt.Println("0~99까지의 숫자중 맞추세요!")
	for {
		fmt.Println("숫자를 입력하세요")
		n, err := InputInt()
		if err != nil {
			fmt.Println("|||||||||||||||||숫자만 입력해주세요|||||||||")
			cnt--
		} else {
			if n > r {
				fmt.Print(n, "보다 작습니다")
			} else if n < r {
				fmt.Print(n, "보다 큽니다")
			} else {
				fmt.Println("숫자를 맞췄습니다 시도횟수:n", cnt, "\t입력숫자:", n)

			}
		}
		fmt.Println("\t현재 시도횟수", cnt)
		cnt++
	}
}

var stdin = bufio.NewReader(os.Stdin)

func InputInt() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
