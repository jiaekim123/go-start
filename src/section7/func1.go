package main

import (
	"fmt"
	"strconv"
)

// 함수 기초 (1)

// 함수 선언 위치는 어느 곳이든 가능
func helloGolang() {
	fmt.Println("ex1 : hello golang")
}

func sayMessage(message string) {
	fmt.Println("ex2 : ", message)
}

func sum(a int, b int) int {
	return a + b
}

func main() {
	// 함수
	// 선언 : func 키워드로 선언
	// func 함수명 (매개변수) (반환 타입 or 반환 값 변수명) : 반환 값 존재
	// func 함수명 () (반환 타입 or 반환 값 변수명) : 매개 변수 없음, 반환 값 존재
	// func 함수명 (매개변수) : 매개변수 존재, 반환 값 없음
	// func 함수명() : 매개변수 없음, 반환 값 없음
	// 타 언어와 달리 반환 값(return value) 다수 가능

	helloGolang()
	sayMessage("hello")
	sayMessage(strconv.Itoa(sum(3, 2))) // 숫자를 문자열로 변환

}
