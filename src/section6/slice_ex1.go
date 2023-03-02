package main

import "fmt"

// 자료형 : 슬라이스 심화(1)

func main() {
	// 슬라이스 추가 및 병합

	// 예제1
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{12, 13, 14, 15, 16}
	s3 := []int{17, 18, 19, 20, 21}
	fmt.Printf("ex1: [len] %d, [cap] %d\n", len(s1), cap(s1))

	s1 = append(s1, 6, 7, 8) // 길이, 용량 값 같이 증가
	fmt.Printf("ex1: [len] %d, [cap] %d\n", len(s1), cap(s1))

	s1 = append(s1, 9, 10, 11) // 길이 증가, 용량 초과 시 2배로 새로 할당
	fmt.Printf("ex1: [len] %d, [cap] %d\n", len(s1), cap(s1))

	s2 = append(s1, s2...)      // 할당된 슬라이스를 append
	s3 = append(s2, s3[0:3]...) // 추출 후 병합

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	// 예제2 - 길이 및 용량 자동 증가
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
	}
}
