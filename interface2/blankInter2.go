package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Lv int
}

func (m *Monster) Attack() {
	fmt.Printf("Monseter Attack")
}

type Player struct {
	Lv int
}

func (m *Player) Attack() {
	fmt.Printf("Player Attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()

		var monster *Monster
		monster = att.(*Monster)
		fmt.Println(monster.Lv)
	}
}

func main() {
	DoAttack(&Monster{20})
}
