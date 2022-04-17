package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6}
	slice1 := append(slice, 9)

	idx := 2
	for i := len(slice1) - 2; i >= idx; i-- {
		slice1[i+1] = slice1[i]
	}
	slice1[idx] = 200
	fmt.Println(slice1)

	slice2 := append(slice[:idx], append([]int{200}, slice[idx:]...)...)
	fmt.Println(slice2)

	slice3 := append(slice, 0)
	copy(slice3[idx+1:], slice[idx:])
	slice3[idx] = 200
	fmt.Println(slice3)
}
