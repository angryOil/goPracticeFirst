package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["oil"] = "서울 성북구"
	m["회사"] = "서울 강남구"
	m["경민"] = "경기도 의정부"

	m["oil"] = "청주시 막걸리"
	fmt.Printf("회사 위치는 %s입니다\n", m["회사"])
	fmt.Printf("oil 은 %s 를 좋아합니다", m["oil"])
}
