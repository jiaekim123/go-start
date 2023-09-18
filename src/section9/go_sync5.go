package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 고루틴 동기화 기초(5)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, notify, notifyAll : 기타 언어
	// Wait, Signal(하나씩 깨우기), Broadcast(전체 깨우기) -> Go언어

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting : ", n)
			condition.Wait() // 고루틴 대기(멈춤)
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		//<-c
		fmt.Println("received : ", <-c)
	}

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()
	//for i := 0; i < 5; i++ {
	//	mutex.Lock()
	//	fmt.Println("Wake Goroutine(Signal) : ", i)
	//	condition.Signal() // 한 개씩 깨움 (모든 고루틴 생성 후)
	//	mutex.Unlock()
	//}

	time.Sleep(1 * time.Second)
}
