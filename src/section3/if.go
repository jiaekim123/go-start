package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF문: 반드시 boolean 검사 -> 1, 0으로 검사하는 것은 안됨! 명시적으로 boolean이어야 함. (자동형변환 불가)
	// 소괄호 사용 x
	var a = 20
	b := 20
	if a >= 15 {
		fmt.Println("a는 15이상")
	}

	if b >= 25 {
		fmt.Println("b는 25이상")
	}

	//// 에러 발생
	//if b >= 25
	//{
	//
	//}

	//if b>= 25
	//	fmt.Println("25이상")
	//if c:= 1; c {
	//	fmt.Println("True")
	//}

	if c := true; c {
		fmt.Println("True")
	}

	//if c:= 40; c>= 35 {
	//	fmt.Println("35이상")
	//}
	//
	//c += 20
}
