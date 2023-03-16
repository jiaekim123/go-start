package main

import "fmt"

// 함수 Closer(1)
func main() {
	// 클로저(Closer)
	// 익명함수를 사용할 경우 함수를 변수에 할당해서 사용 가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부함수에 선언 된 변수에 접근 가능
	// 함수가 선언되는 순간에 함수가 실행될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)

	// 왜 클로저를 쓰는가?
	// 함수를 호출 할 때 이전에 존재했던 값을 유지하기 위해서 사용.
	// -> 비동기, 누적 카운트, 무분별한 전역변수를 줄일 수 있음.
	// 남발하면 객체들이 메모리에 존재하므로, 메모리 부족이나 오버플로우 현상 발생할 수 있음.
	// 클로저를 정확하게 이해하고 사용해야 함.

	// 예제1
	// 클로저를 안쓰는 케이스
	multiply := func(x, y int) int { // 익명 함수
		return x * y // 가비지 컬렉터에 의해 함수가 끝나면 소멸됨.
	}
	r1 := multiply(5, 10)
	fmt.Println("ex1 : ", r1)

	// 예제2
	// 클로저를 사용하는 케이스
	m, n := 5, 10
	sum := func(c int) int { // 함수가 생성되는 순간 변수가 캡쳐됨. 익명함수 변수 할당
		return m + n + c // m, n은 외부에 선언됨.
		// 지역 변수는 소멸하지 않는다. (함수 호출시 마다 사용 가능)
	}

	r2 := sum(10)
	fmt.Println("ex2 : ", r2)

}
