package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}

	sli := arr1[1:2]

	fmt.Println("arr1:", arr1)
	fmt.Println("sli:", sli, len(sli), cap(sli))

	sli = append(sli, 2)
	fmt.Println(sli)
	fmt.Println(arr1)
}
