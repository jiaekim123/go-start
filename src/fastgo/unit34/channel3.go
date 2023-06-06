package main

/*
	채널 버퍼링
	고루틴:  0 고루틴:  1 고루틴:  2
	메인 함수:  0 메인 함수:  1 메인 함수:  2 메인 함수:  3
	위 순서로 실행됨.
*/
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4                 // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true // 채널에 true를 보냄, 버퍼가 가득 차면 대기 // 반복문의 변수 출력
			fmt.Println("고루틴: ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; i < count; i++ {
		<-done                    // 버퍼에 값이 없으면 대기, 값을 꺼냄
		fmt.Println("메인 함수: ", i) // 반복문의 변수 출력
	}
}
