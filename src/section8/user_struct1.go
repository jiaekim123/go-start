package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

// 일반 메소드
func getPrice(c Car) int64 {
	return c.price + c.tax
}

// 구조체 <-> 메소드 바인딩
func (c Car) getPrice() int64 {
	return c.price + c.tax
}

// 사용자 정의 타입 (1)
func main() {
	// Go -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
	// 객체 지향 -> 클래스 (속성: 멤버 변수, 기능(상태: 메소드)) : 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그래밍
	// 객체지향 언어일까? -> 객체지향임.
	// Go는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스 형태의 코딩 가능
	// 객체 지향의 기본 개념 -> Go 에서 포함하고 있다. -> 객체 지향 프로그래밍 언어
	// 상태, 메소드 분리해서 정의(결합성 없음)
	// 사용자 정의 타입: 구조체, 인터페이스, 기본 타입(int, float, string...), 함수
	// 구조체와 -> 메소드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능 (객체 지향)

	// 예제1
	bmw := Car{name: "520d", price: 500000, color: "white", tax: 50000}
	bmwz := Car{name: "220d", price: 600000, color: "white", tax: 60000}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex1 : ", bmwz, &bmwz)

	// 예제2
	// 이건 일반 메소드를 사용하는거임. 객체 지향 X
	fmt.Println("ex2 : ", getPrice(bmw))
	fmt.Println("ex2 : ", getPrice(bmwz))

	// 예제3
	fmt.Println("ex3 : ", bmw.getPrice())
	fmt.Println("ex3 : ", bmwz.getPrice())

	// 예제4
	fmt.Println("ex4 : ", &bmw == &bmwz)

}
