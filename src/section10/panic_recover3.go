package main

import (
	"fmt"
	_ "log"
)

// Panic & Recover(3)
func main() {
	// Go recover 함수
	// 에러 복구 가능
	// Panic으로 발생한 에러를 복구 후 코드 채 실행 (프로그램 종료되지 않는다.)
	// 즉, 코드 흐름을 정상 상태로 복구하는 기능
	// Panicd에서 설정한 메시지를 받아올 수 있다

	// 예제
	runFunc()
	fmt.Println("Hello Golang!")
}

func runFunc() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}
	}()

	a := [3]int{1, 2, 3}
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", a[i]) // 에러 발생 (인덱스 범위 초과)
	}
}
