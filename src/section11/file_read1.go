package main

import (
	"fmt"
	"os"
)

// 파일 읽기
func main() {
	// 파일 읽기
	// Open : 기존 파일 읽기
	// Close : 리소스 열기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 궈한 주의
	// 예외 처리 정말 중요!
	// 탐색 Seek 중요
	// https://golang.org/pkg/os

	// 파일 읽기 예제

	// 파일 열기
	file, err := os.Open("src/section11/sample.txt")
	defer file.Close()
	checkError(err)

	// 읽기 예제1
	fileInfo, err := file.Stat() // 파일 사이즈 확인을 위해서 정보 획득
	checkError(err)
	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담기
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) : ", fileInfo)
	fmt.Println("파일 이름(2) : ", fileInfo.Name())
	fmt.Println("파일 크기(3) : ", fileInfo.Size())
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime())
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n : ", ct1)
	fmt.Println(string(fd1))

	fmt.Println("------------------------------------------")

	// 읽기 예제2 (탐색 seek - offset)
	// offset : 20만큼 커서를 이동하고 읽어온다.
	// whence 옵션 : 0은 처음 위치, 1: 현재 위치, 2 : 마지막 위치
	o1, err := file.Seek(20, 0)
	checkError(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	checkError(err)

	fmt.Printf("읽기 작업(2) 완료 (%d bytes) (%d ret)\n\n", ct2, o1) // ct2 : 읽은 길이, o1 : offset
	fmt.Println(string(fd2))

	// 읽기 예제3
	o2, err := file.Seek(0, 0)
	checkError(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치부터 읽어온다.
	checkError(err)

	fmt.Printf("읽기 작업(3) 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
