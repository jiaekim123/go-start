package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Caculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee     // is a 관계 (상속처럼 사용 가능 - 메소드 재사용)
	specialBonus float64
}

func main() {

	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재사용하려는 장점 제공
	// 상속을 허용하지 않는 Go언어에서 메서드 상속 활용을 위한 패턴

	// 예지1
	// 직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	// 임원
	ep3 := Executives{
		Employee{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Caculate()))
	fmt.Println("ex1 : ", int(ep2.Caculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex1 : ", int(ep3.Caculate()+ep3.specialBonus))
}
