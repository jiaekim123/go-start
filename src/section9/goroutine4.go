package main

import (
	"fmt"
	"runtime"
	"time"
)

// 고루틴(Goroutine) 기초(4)

func main() {
	// 고루틴(Goroutine)
	// 클로저 사용 예제

	// 예제1
	runtime.GOMAXPROCS(1)

	s := "Goroutine Closer : "
	for i := 0; i < 1000; i++ {
		func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) // 반복문 클로저는 일반적으로 즉시 실행 But goroutine으로 한 클로저는 가장 나중에 실행됨. (반복문 종료 후에)
	}

	//time.Sleep(1 * time.Second)
}
