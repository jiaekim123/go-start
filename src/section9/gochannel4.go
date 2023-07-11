package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 채널(Channel)
	// 예제1 (동기: 버퍼 사용 / 비동기)
	// 버퍼: 발신 -> 가득차면 대기함, 비어있으면 다시 작동함.
	// 		수신 -> 비어있으면 대기함, 가득차있으면 작동함.

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2) // 채널의 버퍼 용량을 보냄. (버퍼가 가득 찰때까지 대기함.)
	cnt := 12

	// 클로저를 고루틴으로 실행
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true // 수신할때까지 대기함. (동기라서 하나씩만 보낼 수 있음.)
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
