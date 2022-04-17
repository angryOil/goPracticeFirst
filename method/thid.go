package main

import "fmt"

type account2 struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *account2) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account2) withdrawValue(amount int) account2 {
	a2.balance -= amount
	return a2
}

func main() {
	var mainA *account2 = &account2{100, "j", "oy"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	*mainA = mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

}
