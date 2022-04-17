package main

import "fmt"

type myInt int

func (m myInt) Add(a int) myInt {
	return myInt(int(m) + a)
}

func addFunc(m myInt, a int) myInt {
	return myInt(int(m) + a)
}

func main() {
	var a myInt
	a = a.Add(10)
	fmt.Println(a)
}
