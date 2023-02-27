package main

import "fmt"

// 자료형 : 배열 (2)

func main() {
	// 배열 순회

	// 예제1
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	// Len 길이 반복 (고전적인 방법)
	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1 : ", arr1[i])
	}
	fmt.Println()

	// 예제2
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 { // 이 방식을 더많이 씀
		fmt.Println("ex2 : ", i, v)
	}

	// 인덱스 생략1
	for _, v := range arr2 {
		fmt.Println("ex3 : ", v)
	}

	// 인덱스 생략2
	fmt.Println()
	for v := range arr2 {
		fmt.Println("ex4 : ", v) // 인덱스가 출력됨.
	}
	for i := range arr2 {
		fmt.Println("ex5 : ", arr2[i]) // 값이 필요하면 이렇게.
	}

}
