package main

import "fmt"

// 자료형 : 배열 (1)

func main() {
	// 배열
	// 배열은 용량, 길이 항상 같다
	// 배열 vs 슬라이스 차이점 중요
	// 길이 고정 vs  길이 가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달
	// 전체 비교연산자 사용 가능 vs 비교 연산자 사용불가
	// 대부분 슬라이스 사용한다.

	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 개수

	// 예제1 : 배열
	var arr1 [5]int                         // 기본 값 0
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // 데이터 타입은 모두 일치해야 함.
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}         // 나머지는 0으로 초기화
	arr6 := [...]int{1, 2, 3, 4, 5} // 배열 크기 자동 맞춤 - 잘 사용하지 않는 문법
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	arr1[2] = 5

	// 자료형 & 데이터 타입, 길이, 실제 값 출력
	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr7), arr7)

	// 예제2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "pack"}
	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

}
