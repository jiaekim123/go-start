package main

import "fmt"

// 함수 심화(1)

func main() {
	// 함수 고급
	// 가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	// 예제1
	x := multiplyChangeableParam(1, 2, 3, 4)
	y := sumChangeableParam(1, 23, 3)
	fmt.Println("ex1 : ", x)
	fmt.Println("ex1 : ", y)

	printWord("a", "apple", "test", "seoul")

	// 예제3
	a := []int{1, 2, 3, 4, 5}
	m := multiplyChangeableParam(a...) // 배열을 가변으로 넘기는 방법
	n := sumChangeableParam(a...)

	fmt.Println(m)
	fmt.Println(n)

}

// 가변 파라미터
func multiplyChangeableParam(n ...int) int {
	total := 1
	for _, value := range n {
		total *= value
	}
	return total
}

func sumChangeableParam(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func printWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2: ", value)
	}
}
