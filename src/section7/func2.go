package main

import "fmt"

// 함수 기초 (2)

func calculate(a int, b int, function func(int, int)) {
	function(a, b)
}

func add(a, b int) {
	fmt.Println("ex1 : ", a+b)
}

func minus(a, b int) {
	fmt.Println("ex1 : ", a-b)
}

func multiValue(a int) {
	a *= a
}

func multiReference(a *int) {
	*a *= *a
}

func main() {
	// 함수 (콜백) 전달, 참조 전달 (call by reference), 값 전달(call by value)
	// 예제1
	calculate(100, 110, add)
	calculate(100, 110, minus)

	// 예제2
	// call by value
	a := 100
	multiValue(a)
	fmt.Println(a) // a가 곱해지지 않고 그대로 나옴. (주소 값이 아니므로)

	// 예제 3
	multiReference(&a)
	fmt.Println(a)
}
