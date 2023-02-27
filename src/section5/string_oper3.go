package main

import (
	"fmt"
	"strings"
)

// https://go.dev/doc/
// 데이터 타입 : 문자열 연산(3)

func main() {

	// 예제1 (결함 : 일반 연산) -> 이것도 때에따라 사용
	str1 := "The Go programming language is an open source project to make programmers more productive." +
		"Go is expressive, concise, clean, and efficient. " +
		"Its concurrency mechanisms make it easy to write programs " +
		"that get the most out of multicore and networked machines, " +
		"while its novel type system enables flexible and modular program construction."

	str2 := "It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language."

	// string은 변수를 수정하는게 아니라 새로운 변수를 만들어서 넣는거기 때문에 메모리나 성능상 안좋다는 말이 있으나 많이 사용하는 방식.
	fmt.Println("ex1 : ", str1+str2)

	// 예제2 (결합 : Join) -> 일반적으로 추천되는 방식
	// 불필요한 연산을 최소화할 수 있음.
	// 결합이 많이 있을 경우 join을 사용하는 것이 좋음. (썽능상 훨씬 이점이 있음.)
	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, "----"))

}
