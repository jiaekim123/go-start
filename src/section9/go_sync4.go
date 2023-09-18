package main

// 고루틴 동기화 기초(4)

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 뮤텍스: 상호 배제
	// Thread(고루틴)들이 서로 running time에 서로 영향을 주지 않게 즉, 단독으로 실행되게 하는 기술
	// 뮤텍스: 여러 고루틴에서 작업 공유 데이터를 보호한다.
	// RWMutex : 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 X, 읽기 락, 쓰기 락 모두 방어
	// RMutex : 읽기 Lock -> 읽기 시도 중에 값이 변경 방지, 즉 쓰기 락 방어(방지)

	// 동기화 사용하는 경우
	// 읽기, 쓰기 전용 뮤텍스를 사용하는 경우
	// 쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환할 가능성을 없애줌.

	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())
	data := 0
	mutex := new(sync.RWMutex) // var mutex = new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			// 쓰기 뮤텍스 잠금 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 :", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
