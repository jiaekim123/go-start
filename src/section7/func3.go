package main

import "fmt"

// 함수 기초 (3)

// 값 여러개 반환 가능
func multiply(x int, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	// 다중 값 반환
	// 예제1
	a, b := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)
	fmt.Println("ex1 : ", a, b, c, d)

	// 예제2
	x1, x2, x3, x4, x5 := arrayMultiply(10, 20, 30, 40, 50)
	y1, _, _, y4, _ := arrayMultiply(10, 20, 30, 40, 50)
	fmt.Println("ex2 : ", x1, x2, x3, x4, x5)
	fmt.Println("ex2 : ", y1, y4)
}
