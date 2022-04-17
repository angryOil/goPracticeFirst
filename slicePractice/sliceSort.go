package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	slice := []int{5, 4, 2, 1, 3}
	sort.Ints(slice)
	fmt.Println(slice)

	s := []Student{
		{"joy", 51},
		{"world", 33},
		{"warm", 13},
		{"oil", 1},
		{"angry", 15},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}
