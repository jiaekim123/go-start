package main

import "fmt"

// 슬라이스 심화(3)

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사 해야 한다.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 없음.

	// 예제1(복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)

	// 예제2

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)
	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a) // 원본에는 영향이 가지 않았음!
	fmt.Println("ex2 : ", b)

	fmt.Println()

	// 예제3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]

	d[1] = 7 // 주의: 부분적으로 슬라이스 추출은 참조 -> 원본 값 변경됨.
	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	fmt.Println()

	// 예제4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5]   // 슬라이스 참조 값 복사해옴.
	g := e[0:5:7] // 마지막꺼는 capacity 지정

	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)

	fmt.Println("ex4 : ", len(g), cap(g))
	fmt.Println("ex4 : ", g)
}
