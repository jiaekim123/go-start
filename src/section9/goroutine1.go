package main

import (
	"fmt"
	"time"
)

// 고루틴(Goroutine) 기초(1)

func exe1() {
	fmt.Println("ex1 function start : ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ex1 function end : ", time.Now())
}

func exe2() {
	fmt.Println("ex2 function start : ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ex2 function end : ", time.Now())
}

func exe3() {
	fmt.Println("ex3 function start : ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ex3 function end : ", time.Now())
}

func main() {
	// 고루틴(Goroutine)
	// 타 언어의 쓰레드(Thread)와 비슷한 기능
	// 생성 방법 매우 간단, 리소스 매우 적게 사용
	// 즉, 수많은 고루틴 동시 생성 실행 가능
	// 비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통한 통신 가능
	// 공유 메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글루틴에 비해 항상 빠른 처리 결과는 아님.

	// 멀티 쓰레드 장점과 단점
	// 장점: 응답성 향상, 자원공유를 효율적으로 사용, 작업이 분리되어 코드가 간결해짐
	// 단점: 구현하기 어려움, 테스트 및 디버깅이 어려움, 전체 프로세스의 사이드 이펙트, 성능 저하,
	//		동기화 코딩 반드시 숙지, 데드락 등
	// -> 정확하게 알고 써야 한다!
	// 파이썬 쓰레드 추천! (입문용) => 이론을 예제를 통해 꼭 학습해보기 (언어가 뭐가 되었든 개념을 정확하게 파악하기)

	exe1() // 가장 먼저 실행 (일반적인 실행 흐름)
	// 메인 쓰레드가 끝남과 동시에 다른 쓰레드들이 다 같이 끝나는 것 => 데몬 쓰레드

	fmt.Println("Main Routine Start", time.Now())

	go exe2()
	go exe3()

	//time.Sleep(4 * time.Second)
	fmt.Println("Main Routine End", time.Now())
}
