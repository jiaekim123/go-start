package main

import "fmt"

// 채널(Channel) 기초 (5)

func main() {

	// 채널(Channel)
	// Close : 채널 닫기

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good!"
		}
	}()

	val1, ok1 := <-ch // 채널이 정상이면 ok1는 true가 나옴.
	fmt.Println("ex1 : ", val1, ok1)

	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)

	close(ch)         // 채널 닫기
	val3, ok3 := <-ch // val3은 기본값인 빈문자열임.
	fmt.Println("ex1 : ", val3, ok3)

}
