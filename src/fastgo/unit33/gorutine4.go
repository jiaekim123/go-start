package main

/*
	33.2 클로저를 고루틴으로 실행하기
*/

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // 싱글코어 사용

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

	s := "Hello, hello!"

	for i := 0; i < 100; i++ {
		go func(n int) { // 익명 함수를 고루틴으로 실행
			fmt.Println(s, n)
		}(i)
	}
	fmt.Scanln()
}
