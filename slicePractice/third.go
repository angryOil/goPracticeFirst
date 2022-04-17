package main

import "fmt"

func addNum(slice []int) []int {
	slice = append(slice, 4)
	fmt.Println(slice)
	return slice
}
func addNum2(slice *[]int) {
	*slice = append(*slice, 4)
	fmt.Println(slice)
}

func main() {
	slice := []int{1, 2, 3, 1}
	addNum2(&slice)
	slice = addNum(slice)
	fmt.Println(slice)
}
