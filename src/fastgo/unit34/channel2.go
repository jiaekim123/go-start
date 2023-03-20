package main

/*
	34.1 동기 채널
	고루틴에서 채널을 통해 done을 보내야 메인 함수가 실행됨.
	고루틴 0 -> 메인함수 0 -> 고루틴 1 -> 메인함수 1 -> 고루틴 2 -> 메인함수 2 순서로 실행됨.
	동기 채널을 사용하면 고루틴의 코드 실행 순서를 제어할 수 있음.
*/
import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // 동기 채널 샛엉
	count := 3              // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // 고루틴에 true를 보냄, 값을 꺼낼때까지 대기
			fmt.Println("고루틴 : ", i)    // 반복문의 변수 출력
			time.Sleep(1 * time.Second) // 1초 대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done                     // 채널에 값이 들어올 때까지 대기, 값을 꺼냄
		fmt.Println("메인 함수 : ", i) // 반복문의 변수 출력
	}
}
