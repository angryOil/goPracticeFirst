package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	slice1 := make([]int, len(slice))
	copy(slice1, slice)
	idx := 2
	for i := idx + 1; i < len(slice1); i++ {
		slice1[i-1] = slice1[i]
	}
	slice1 = slice1[:len(slice1)-1]
	fmt.Println(slice1)

	slice2 := append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice2)
}
