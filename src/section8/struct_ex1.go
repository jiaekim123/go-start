package main

import "fmt"

type MyAccount struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *MyAccount {
	// 포인터 반환이 아니면 항상 값을 복사한다.
	// 구조체 인스턴스를 생성한 뒤 포인터를 반환한다.
	return &MyAccount{number, balance, interest}
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제1
	// 생성자를 통해 멤버 변수를 초기화
	kim := MyAccount{number: "245-901", balance: 100000, interest: 0.015}

	var lee *MyAccount = new(MyAccount)
	lee.number = "245-902" // getter, setter
	lee.balance = 130000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	// 예제2
	// 구조체를 반환할 때 무조건 '포인터 반환'을 하고 싶다면 생성자 패턴을 통해 사용한다.
	park := NewAccount("245-903", 170000, 0.04)
	fmt.Println("ex1 : ", park)
}
