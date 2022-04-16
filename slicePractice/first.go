package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4}
	slice2 := append(slice, 4) //기존 슬라이스는 그대로고 추가된값을 반환

	fmt.Println(slice)
	fmt.Println(slice2)

}
