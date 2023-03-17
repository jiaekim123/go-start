package main

import "fmt"

// 사용자 정의 타입 (2)

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입
	// 예제1
	a := cnt(15)
	fmt.Println("ex1 : ", a)

	// 예제2
	var b cnt = 15 // 보통은 이렇게 사용. (사용자 정의 타입)
	fmt.Println("ex1 : ", b)

	// 예제3
	//testConvertT(b) // 기본 자료로 만들었으면 cnt라는 사용자 정의 타입은 못넣음. (에러)
	// 예외 발생(중요)! 사용자 정의 타입 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용 가능함.
	testConvertT(int(b)) // 직접 형변환 해서 넣는 것은 가능.
	testConvertD(b)

}

func testConvertT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex3 : (Custom Type) : ", i)
}
