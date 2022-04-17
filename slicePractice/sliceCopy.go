package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}
	slice3 := append([]int{}, slice1...)
	slice4 := make([]int, len(slice1))
	copy(slice4, slice1)
	slice2[2] = 200
	slice3[3] = 11
	slice4[2] = 4
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
