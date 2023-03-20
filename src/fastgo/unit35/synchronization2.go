package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	34.1 뮤텍스로 데이터 보호하기
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() { // 고루틴에서
		for i := 0; i < 1000; i++ {
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // data 슬레이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))
}
