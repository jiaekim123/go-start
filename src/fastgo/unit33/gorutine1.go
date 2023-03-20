package main

import "fmt"

func hello() {
	fmt.Println("Hello world!")
}
func main() {
	go hello() // 함수를 고루틴으로 실행

	fmt.Scanln() // 종료되지 않도록 실행
}
