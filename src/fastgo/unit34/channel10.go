package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) // int형 채널 생성

	go func() {
		for {
			i := <-c1                          // 채널 c1에서 값을 꺼낸 후 i에 대입
			fmt.Println("c1 : ", i)            // i 값을 출력
			time.Sleep(100 * time.Millisecond) // 100밀리초 대기
		}
	}()

	go func() {
		for {
			c1 <- 20                           // 채널 c1에 20을 보낸 뒤
			time.Sleep(500 * time.Millisecond) // 100밀리초 대기
		}
	}()

	go func() {
		for {
			select {
			case c1 <- 10: // 매번 채널 c1에 10을 보냄
			case i := <-c1: // c1에 값이 들어왔을 때는 값을 꺼낸 뒤 i에 대입
				fmt.Println("c1 : ", i) // i값을 출력
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10초동안 프로그램 실행
}
