package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	대기 그룹 사용하기
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1) // 반복될 때마다 wg.Add 함수로 1씩 추가
		go func(n int) {
			fmt.Println(n)
			wg.Done() // 고루틴이 끝났다는 것을 알려줌
		}(i)
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다림
	fmt.Println("the end")
}
