package main

import "fmt"

// 데이터 타입 : Bool(1)

func main() {
	// Bool(Boolean): 참과 짓짓
	// 조건무 논리 연산자랑 주로 사용 : !(not), ||(or), && (and)
	// 암묵적 형변환x : 0, Nil -> False 취급 안함.

	// 예제1
	var b1 bool = true
	var b2 bool = false
	b3 := true
	fmt.Println("ex1 : ", b1, b2, b3)

	fmt.Println("ex2 : ", b1 == b3)
	fmt.Println("ex3 : ", b1 && b2)
	fmt.Println("ex4 : ", b1 || b2)
	fmt.Println("ex5 : ", !b1)

	//b4 := 1
	//if b4 {
	//	fmt.Println("ex6 : true")
	//}
}
