package main

import "fmt"

// 자료형 : 포인터(1)

func main() {
	// 포인터
	// Go : 포인터 지원(C)
	// 변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택
	// 파이썬, 자바(JRE) -> 컴파일러, 인터프리터
	// 포인터 지원 X(파이썬, C#, JAVA 등)
	// 주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
	// *(에스터 리스크)로 사용
	// nil로 초기화 (nil != 0)

	// 예제1
	//var a *int = 5 // 주소값을 넣어야 함.
	var a *int            // 방법1
	var b *int = new(int) // 방법2

	fmt.Println(a)
	fmt.Println(b)

	i := 7

	a = &i
	b = &i

	fmt.Println("ex1 : ", a, &i) // i값에 대한 번지수가 저장
	fmt.Println("ex1 : ", &a)    // a에 대한 주소 값
	fmt.Println("ex1 : ", *a)    // 역참조

	fmt.Println()
	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b)

	var c = &i
	d := &i

	*d = 10 // 역참조한 값을 변경함.

	fmt.Println("ex1 : ", c, &i) // i값에 대한 번지수가 저장
	fmt.Println("ex1 : ", &c)    // a에 대한 주소 값
	fmt.Println("ex1 : ", *c)    // 역참조

	fmt.Println()
	fmt.Println("ex1 : ", d, &i)
	fmt.Println("ex1 : ", &d)
	fmt.Println("ex1 : ", *d)

}
