// go 특징 첫번째 강의
package main

import "fmt"

func main() {
	// go: 모호하거나, 애매한 문법 금지

	// 후위 연산자 허용 i++ (o)
	// 전위 연산자 비허용 ++i (x)
	// 증감 연산 반환 값 없음. sum i+ i++ -> 반환값 없으므로 이렇게 사용 불가

	// 포인터(Pointer) - go에서 포인터 허용. 단, 연산은 비허용.

	// 주석 (// /**/)

	// 예제1
	sum, i := 0, 0
	for i <= 100 {
		//sum += i++ // 안됨!
		sum += i // 가능
		i++
		// ++i // 예외 발생
	}
	fmt.Println("ex1: ", sum)
}
