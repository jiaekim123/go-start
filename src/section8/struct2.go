package main

import "fmt"

// 구조체 기본(1)

// Account 계좌
type account struct {
	number   string
	balance  float64 // 잔액
	interest float64 // 이자
}

func (a account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 다양한 선언 방법
	// &struct, struct : 이 둘의 차이점은 &struct는 포인터를 받아오고 역참조를 하기 때문에 속도는 조금 더 느림.
	// 다만 &struct로 사용해야 할 경우가 있는데, Go 인터페이스에서 메소드를 선언만 해둔 후
	//-> overriding해서 메서드에 포인터 리시버를 사용할 경우 반드시 사용
	//-> 이 케이스에는  new(Account)로 넘겨주어야 함. (포인터 리시버 형태로)

	// 선언 방법1
	var kim *account = new(account) // new를 쓸 때에는 포인터를 써야 함.
	kim.number = "245-901"
	kim.balance = 1000000
	kim.interest = 0.015

	// 선언 방법2
	hong := &account{number: "245-902", balance: 1500000, interest: 0.04}

	// 선언 방법3
	lee := new(account) // new는 선언만할 수 있음.
	lee.number = "245-903"
	lee.balance = 1000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)

	fmt.Printf("ex2 : %#v\n", kim) // 전체 값 출력
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)
	fmt.Println()

	// 예제2
	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(hong.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))

}
