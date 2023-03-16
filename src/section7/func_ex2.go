package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}
func main() {
	// 함수 고급
	// 함수를 변수에 할당

	// 예제1 (슬라이스에 할당)
	f := []func(int, int) int{multiply, sum} // 함수를 두개 입력받음

	a := f[0](10, 10) // multiply 함수
	b := f[1](10, 10) // sum 함수

	fmt.Println("ex1 : ", a, f[0](10, 10))
	fmt.Println("ex1 : ", b, f[1](10, 10))

	// 예제2 (변수에 할당)
	var f1 func(int, int) int = multiply
	f2 := sum // 이렇게 많이 사용함.
	fmt.Println("ex2 : ", f1(10, 10))
	fmt.Println("ex2 : ", f2(10, 10))

	// 예제2 (맵에 할당)
	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}
	fmt.Println("ex3 : ", m["mul_func"](10, 10))
	fmt.Println("ex3 : ", m["sum_func"](10, 10))
}
