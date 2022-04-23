package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]string{}
	m[hash(23)] = "warmOil"
	m[hash(255)] = "angryOil"

	fmt.Printf("%d = %s\n", 23, m[hash(23)])
	fmt.Printf("%d = %s\n", 255, m[hash(255)])

}
