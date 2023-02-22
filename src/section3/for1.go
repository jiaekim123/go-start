package main

import "fmt"

func main() {
	// 반복문 - for
	// go에서 유일하게 제공되는 반복문
	// 다양한 사양법 숙지

	// 예제1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	//// 에러 발생 1
	//for i := 0; i < 5; i++ // 이렇게 하면 안됨! go 내부적으로 세미콜론을 찍음
	//{
	//
	//}

	//// 에러 발생2
	//for i := 0; i < 5; i++
	//	fmt.Println("")

	//// 예제 2 (무한루프)
	//for {
	//	fmt.Println("ex2 : Hello, Golang!")
	//	fmt.Println("ex2 : Infinite loop")
	//}

	// 예제3 (Range 비법)
	loc := []string{"seoul", "busan", "inchon"}
	for index, name := range loc {
		fmt.Println("ex3: ", index, name)
	}

	// 예제3 (Range 비법)
	loc2 := []string{"seoul", "busan", "inchon"}
	for _, name := range loc2 {
		fmt.Println("ex3: ", name)
	}
}
