package main

import "fmt"

func main() {
	// defer 함수 실행 (지연)
	// defer를 호출한 함수가 종료되기 직전에 호출
	// 타 언어의 finally 문과 비슷
	// 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제

	// 예제1
	exF1()
}

func exF1() {
	fmt.Println("f1: start!")
	defer exF2() // 마지막에 호출됨.
	fmt.Println("f1: end!")
}

func exF2() {
	fmt.Println("f2: called!")
}
