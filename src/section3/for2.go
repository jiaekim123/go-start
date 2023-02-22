package main

import "fmt"

func main() {
	// 예제1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1
	}
	fmt.Println("ex1: ", sum1)

	// 예제2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		//j := i++ // go에서 후치연산은 반환값이 없음.
	}
	fmt.Println("ex2: ", sum2)

	// 예제3 - while과 비슷한 형태로 사용
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3: ", sum3)

	// 예제4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4: ", i, j)
	}

	// 에러 발생
	//for i, j := 0, 0; i <= 10; i++, j += 10{// 후치연산은 반환값이 없음.
	//
	//}
}
