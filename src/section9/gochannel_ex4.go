package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널(Channel) 셀렉트 구문
	// 채널 Select 구문 -> 채널에 값이 수신되면 해당 case 문 실행
	// 일회성 구문이므로, For문으로 안에서 수행한다.
	// default 구문 처리에 주의한다.

	// 예제1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Microsecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!"
			time.Sleep(500 * time.Microsecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
				//default: // 값이 도착하기 전에 default가 실행되는 케이스가 발생하므로 채널에서는 default를 안씀.
				//	fmt.Println("default test")
			}
		}
	}()
	time.Sleep(7 * time.Second)
}
