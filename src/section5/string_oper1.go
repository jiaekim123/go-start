package main

import "fmt"

// 데이터 타입 : 문자열 연산(1)

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// 예제1 (추출)
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0]) // 슬라이싱을 해서 가져오면 문자열이 출력, 하나만 출력하면 정수형 출력
	fmt.Println("ex1 : ", str2[3:], str2[0])
	fmt.Println("ex1 : ", str2[:4])
	fmt.Println("ex1 : ", str1[1:3])
}
