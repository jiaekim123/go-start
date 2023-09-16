package main

import (
	"errors"
	"fmt"
	"math"
)

// Go 에러 처리 고급 (1)

func main() {
	// 에러 처리 고급
	// Go error 패키지 New 메소드 활용 -> 사용자 에러 처리 생성

	// 예제1
	if f, err := Power(7, 3); err != nil {
		fmt.Printf("Error message : %s\n", err)
		//fmt.Printf("Error message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	// 예제2
	if f, err := Power(0, 3); err != nil {
		fmt.Printf("Error message : %s\n", err)
		//fmt.Printf("Error message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}
}

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용 불가")
	}
	return math.Pow(f, i), nil // 제곱, nil 반환
}
