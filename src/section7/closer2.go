package main

import "fmt"

func main() {
	// 예제1
	cnt := increaseCnt()
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt()) // 함수가 호출될 때 계속 값이 누적됨.
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	// 메모리 누수 주의

}

func increaseCnt() func() int {
	n := 0 // 지역 변수 (캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
