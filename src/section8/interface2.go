package main

import "fmt"

// 인터페이스 기본 (2)

type Dog struct {
	name   string
	weight int
}

// byte 메소드 구현 (리시버)
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

// Animal 동물이 행동하는 인터페이스 선언
type Animal interface {
	bite()
}

func main() {
	// 인터페이스 구현 예제
	// 예제1
	dog1 := Dog{"poll", 10}
	var animal Animal
	animal = dog1
	animal.bite()

	//dog1.bite() // 이게 아니라 interface의 메서드를 호출해도 dog이 실행됨.

	// 예제2
	dog2 := Dog{"marry", 12}
	animal2 := Animal(dog2)
	animal2.bite()

	// 예제3
	animals := []Animal{dog1, dog2}
	// 인덱스 형태로 실행
	for index, _ := range animals {
		animals[index].bite()
	}

	// 값 형태로 실행(인터페이스)
	for _, val := range animals {
		val.bite()
	}
}
