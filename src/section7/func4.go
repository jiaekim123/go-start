package main

import "fmt"

// 함수 기초 (4)

func multiplyWithReturnValueName(x int, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func main() {

	// 예제1
	a, b := multiplyWithReturnValueName(1, 2)
	fmt.Println("ex1 : ", a, b)
}
