package main

import "fmt"

// 채널 심화 3

func main() {
	// 채널(Channel)

	// 예제1
	c := receiveChannelOnly(100) // 채널 반환
	output := total(c)           // 채널 전달 후 반환

	fmt.Println("ex1 : ", <-output)

}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}

func receiveChannelOnly(cnt int) <-chan int {
	sum := 0
	total := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		total <- sum
		total <- 888
		total <- 777
		close(total) // close해야 끝남!
	}()
	return total
}
