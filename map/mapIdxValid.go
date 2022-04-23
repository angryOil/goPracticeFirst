package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 3
	m[3] = 4
	delete(m, 3)
	delete(m, 4)
	v, ok := m[3]
	fmt.Println(v, ok)
	v, ok = m[1]
	fmt.Println(v, ok)
}
