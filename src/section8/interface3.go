package main

import "fmt"

// 인터페이스 기본 (3)

type dog struct {
	name   string
	weight int
}

type cat struct {
	name   string
	weight int
}

// 구조체 Dog 메서드 구현
func (d dog) bite() {
	fmt.Println(d.name, " : Dog bites!")
}

func (d dog) sounds() {
	fmt.Println(d.name, " : Dog barks!")
}

func (d dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

// 구조체 Cat 메서드 구현
func (d cat) bite() {
	fmt.Println(d.name, " : Cat 할퀴다!")
}

func (d cat) sounds() {
	fmt.Println(d.name, " : Cat cries!")
}

func (d cat) run() {
	fmt.Println(d.name, " : Cat is running!")
}

// 메서드만으로 동물인지 판단할 수 있음.
type animal interface {
	bite()
	sounds()
	run()
}

// 인터페이스를 파라미터로 받는다.
func action(ani animal) {
	ani.bite()
	ani.sounds()
	ani.run()
}

func main() {
	// 인터페이스의 구현 예제
	// 인터페이스 규격화 역할 이해
	// 인터페이스에 정의된 메소드 사용 유도
	// 코드의 가독성 및 유지보수 증가

	// 덕타이핑 예제
	// 덕타이핑: 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	// Go의 중요한 특징: 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다. (따로 implements를 받지 않아도 됨.)

	// 에제1
	dog1 := dog{"hello", 10}
	cat1 := cat{"meeu", 11}

	action(dog1)
	action(cat1)
}
