package main

import "fmt"

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int) // 클로저로 캡쳐가 된 상태
	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}
func main() {
	// 채널(Channel)
	// 채널 또한 함수의 반환 갑승로 사용 가능

	// 예제1
	c := sum(100)

	fmt.Println("ex1 : ", <-c)

}
