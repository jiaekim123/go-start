package main

import "fmt"

func main() {
	// 코드 서식 유틸
	// 한 사람이 개발한 것 같은 일관성 유지
	// 코드 스타일 유지
	// gofmt -h : 사용법
	// gofmt -w : 원본 파일에 반영

	// 예제1
	for i := 100; i <= 100; i++ {
		fmt.Println("ex1: ", i)
	}
}
