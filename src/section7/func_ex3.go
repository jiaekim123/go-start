package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func printHello(n int) {
	if n == 0 {
		return
	}

	fmt.Println("ex2 : (", n, ")", "hi!")
	printHello(n - 1)
}

func main() {
	// 함수 고급
	// 재귀 함수(Recursion)
	// 프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이: 장점
	// 코드를 이해하기 어렵고, 기억공간을 많이 사용, 무한루프 가능성

	// 예제1
	x := fact(7)
	fmt.Println("ex1 : ", x)

	// 예제2
	printHello(3)
}
