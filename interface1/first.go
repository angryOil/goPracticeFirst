package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕 나는 %d 살 %s 이라고해", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	s1 := Student{"철인", 22}
	var stringer Stringer
	stringer = s1
	fmt.Println(stringer.String())
	fmt.Println(s1.GetAge())
	fmt.Println(stringer)
}
