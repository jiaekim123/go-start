package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 고루틴(Goroutine) 기초(3)

func exe4(num int) {
	r := rand.Intn(100)

	fmt.Println(num, " start ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(num, " >>>> ", r, i) // r은 랜덤 값
	}
	fmt.Println(num, " end ", time.Now())
}
func main() {
	// 고루틴(Goroutine)
	// 멀티 코어(다중 CPU) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU()) // PC의 최대 CPU 값을 전부 사용
	// CPU가 일할 때 다른 프로세스가 멈추면 CPU를 전부 쓴다고 무조곤 빠른건 아님. (프로그램을 잘자야 빠름)
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 0으로 하면 현재 내가 셋팅해둔 값을 출력함.

	// 예제1
	fmt.Println("Main Routine Start : ", time.Now())

	for i := 0; i < 100; i++ {
		go exe4(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Main Routine End : ", time.Now())

}
