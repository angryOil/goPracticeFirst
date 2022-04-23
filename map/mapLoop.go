package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)
	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 300}
	m[888] = Product{"자", 1000}

	for k, v := range m {
		fmt.Printf("%d 번 품목: %s 가격%d\n", k, v.Name, v.Price)
	}
}
