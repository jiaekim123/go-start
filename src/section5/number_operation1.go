package main

// 데이터 타입: Numeric 연산(1)

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산 (산술, 비교)
	// 타입이 같아야 가능
	// 숫자 계산 시에는 손으로 계산한 뒤 코드로 된 결과를 한번 더 확인 해야 함. (통화에서 사용되는 것은 반드시 확인)
	// 다른 타입끼리는 반드시 형변환 후 연산, 형변환 없을 경우 예외(에러) 발생
	// +, -, *, %, /, <<, >>, &, ^

	// 예제 1
	// math.Max, Min으로 최대 최소 값이 어떤건지 확인 가능
	var n1 uint8 = math.MaxUint8 // Uint는 unsigned int라 그냥 int보다 max값이 두배임.
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1 : ", n1)
	fmt.Println("ex1 : ", n2)
	fmt.Println("ex1 : ", n3)
	fmt.Println("ex1 : ", n4)

	fmt.Println("ex1 : ", math.MinInt8)
	fmt.Println("ex1 : ", math.MinInt16)
	fmt.Println("ex1 : ", math.MinInt32)
	fmt.Println("ex1 : ", math.MinInt64)

	fmt.Println("ex1 : ", math.MaxFloat32)
	fmt.Println("ex1 : ", math.MaxFloat64)

	n5 := 10000        // int
	n6 := int16(10000) // 형변환
	n7 := uint8(100)
	//n7 := uint8(300) // 에러

	//fmt.Println("ex3 : ", n5 + n6) // 다른 타입끼리는 연산 불가
	fmt.Println("ex3 : ", n5+int(n6))
	fmt.Println("ex3 : ", n6+int16(n7))
	fmt.Println("ex2 : ", n6 > int16(n7))
	fmt.Println("ex2 : ", n6-int16(n7) > 5000)

}
