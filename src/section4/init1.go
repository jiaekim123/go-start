package main

import (
	"fmt"
)

func init() {
	// main 함수보다 먼저 실행됨 (1회)
	// 프로그램 초기 변수 설정 or 시작 전 프로그램 상태 확인 등
	// 패키지 로드 시에 가장 먼저 호출됨.
	fmt.Println("init method start!")
}
func main() {

	fmt.Println("main method start!")
}
