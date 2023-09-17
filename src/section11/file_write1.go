package main

import (
	"fmt"
	"os"
)

// 파일 쓰기

func main() {
	// 파일 쓰기
	// Create : 새 파일 작성 및 파일 열기
	// Close : 리소스 닫기
	// Write, WriteString, WriteAt (특정 위치부터 작성) : 파일 쓰기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요!
	// https://golang.org/pkg/os

	// 파일 쓰기 예제
	file, err := os.Create("test_write.txt")
	ErrCheck2(err)

	// 리소스 해제
	defer file.Close()

	// 쓰기 예제 1
	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 형변환 후 쓰기 (생략 가능)
	ErrCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n\n", n1)
	file.Sync() // Write Commit (Stable)

	// 쓰기 예제 2
	s2 := "Hello Golang\n File Write Test - 1\n"
	n2, err := file.WriteString(s2)
	ErrCheck2(err)

	fmt.Printf("쓰기 작업(2) 완료 (%d bytes) \n\n", n2)

	s3 := "Test WriteAt! - 2\n"
	n3, err := file.WriteAt([]byte(s3), 100) // len(offset)
	ErrCheck2(err)
	fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n\n", n3)

	file.Sync()

	// 쓰기 예제4
	n4, err := file.WriteString("Hello Golang! \n File Write Test - 4\n")
	ErrCheck2(err)
	fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n\n", n4)
}

// ErrCheck1 에러 체크 방식 1
//func ErrCheck1(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

// ErrCheck2 에러 체크 방식 2
func ErrCheck2(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
