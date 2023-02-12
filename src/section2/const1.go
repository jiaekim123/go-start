package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 훟 값 변경 금지, 고정 된 값 관리 용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	//const d = getHeight() // 함수로는 초기화 할 수 없음.
	const e = 35.67
	const f = false

	/*
		에러 발생
		const g string
		g = "Test2"
	*/

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	//fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
}
