package main

import "fmt"

// 구조체 기본(1)

// Account 계좌
type Account struct {
	number   string
	balance  float64 // 잔액
	interest float64 // 이자
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 구조체
	// Go의 필드들의 집합체 또는 컨테이너
	// 필드는 갖지만 메소드는 갖지 않는다.
	// 객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	// 상속, 객체, 클래스 개념 없음
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	// 예제1
	kim := Account{number: "245-901", balance: 1000000, interest: 0.03}
	lee := Account{number: "245-902", balance: 2500000} // 기본값 0
	pack := Account{number: "245-903", interest: 0.015} // 기본값 0
	cho := Account{"245-904", 1000000, 0.02}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println("ex1 : ", pack)
	fmt.Println("ex1 : ", cho)
	fmt.Println()

	// 예제2
	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(pack.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))

}
