package main

import (
	"errors"
	"fmt"
	"log"
)

// Go 에러 처리 기초 (2)

func main() {
	// 에러 처리
	// Errorf를 이용한 에러 처리 (사용자 정의 처리)

	a, err := isNotZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", a)

	b, err := isNotZero(0)
	if err != nil {
		log.Fatal(err)
	}

	// 프로그램 종료 후 실행 중지
	fmt.Println("ex2 : ", b)
	fmt.Println("End Error Handling!")
}

func isNotZero(num int) (string, error) {
	if num != 0 {
		s := fmt.Sprint("Hello Golang : ", num)
		return s, nil
	}
	return "", errors.New("0은 입력할 수 없습니다. 에러 발생")
}
