package main

import "fmt"

// 인터페이스 기본 (4)

type galaxy struct {
	name   string
	weight int
}

type ipone struct {
	name   string
	weight int
}

func (d galaxy) on() {
	fmt.Println(d.name, " : 갤럭시가 커졌습니다!")
}

func (d ipone) on() {
	fmt.Println(d.name, " : 아이폰이 켜졌습니다!")
}

func act(phone interface{ run() }) { // 익명 인터페이스 (타입 정의x)
	phone.run()
}

func main() {

	// 익명 인터페이스 사용 예제(즉시 선언 후 사용)

	// 에제1
	gal := dog{"galaxy21", 10}
	ip := cat{"iphone10", 11}

	act(gal)
	act(ip)
}
