package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널(Channel)
	// 예제1 (동기: 버퍼 미사용)

	ch := make(chan bool)
	cnt := 6

	// 클로저를 고루틴으로 실행
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true // 수신할때까지 대기함. (동기라서 하나씩만 보낼 수 있음.)
			fmt.Println("Go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
