package main

import "fmt"

func main() {
	// 문장 끝 세미콜론(;) 주의
	// 자동으로 끝에 go 컴파일러가 끝에 세미콜론을 삽입함.
	// 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용 (권장하지 않음.)

	// 반복문 및 제어문 (조건문) (if, for) 사용할 경우에 주의

	// 예제1
	for i := 0; i <= 10; i++ {
		//fmt.Println("ex1 : ", i); fmt.Println(i)
		fmt.Print("ex1 : ", i)
		fmt.Println(i)
	}
	fmt.Println()

	// 예제2
	for j := 10; j >= 0; j-- {
		fmt.Println("ex2: ", j)
	}
}
