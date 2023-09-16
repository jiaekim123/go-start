package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// Go 에러 처리 고급(3)

// 예외(에러) 처리 구조체
type PowError struct {
	time time.Time // 에러 발생 시간
	//value interface{}
	value   float64 // 파라미터
	message string  // 에러 지시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v] Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

func power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) || math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: f, message: "슷지만 사용해야 합니다."}
	}
	return math.Pow(f, i), nil
}

func main() {
	// 에러 처리 고급
	// error 타입이 아닌 경우 에러 처리 방법
	// Error 메소드 구현해서 사용자 젖ㅇ의 에러 처리 예제 심화
	// 구조체를 사용해서 세부적인 정보 출력

	// 예제1
	v, err := power(10, 3) // 정상 처리
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ex1 : ", v)

	// 예제2
	t, err := power(0, 3) // 예외 발생
	if err != nil {
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
		fmt.Println(err.(PowError).time)
		log.Fatal(err.Error())
	}
	fmt.Println("ex1 : ", t)
}
