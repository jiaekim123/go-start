package main

import "fmt"

// 채널(Channel) 기초(2)

func main() {
	// 채널(Channel)

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700, c)
	go rangeSum(500, c)

	// 순서대로 데이터 수신(동기) : 채널에서 값 수신 완료될때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("ex1 : ", result1)
	fmt.Println("ex1 : ", result2)
	fmt.Println("ex1 : ", result3)
}

func rangeSum(rg int, c chan int) {
	sum := 0

	// 0부터 rg까지 더하고 채널로 보낸다.
	for i := 0; i <= rg; i++ {
		sum += i
	}
	c <- sum
}
