package main

import "fmt"

func main() {
	// 예제 1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2, 4, 6인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // i가 1, 3, 5인 경우
		fmt.Println("a -> ", a, "는 홀수")
	}

	// 예제2
	switch e := "go"; e {
	case "java":
		fmt.Println("java")
	case "go":
		fmt.Println("go")
	case "python":
		fmt.Println("python") // 실행 안함 (자동 break)
	case "ruby":
		fmt.Println("ruby") // 실행 안함 (자동 break)
	case "c":
		fmt.Println("c") // 실행 안함 (자동 break)
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough // 반드시 다음에 실행되어야 하는 것이 있으면 출력 (break를 안쓰겠다는 뜻)
	case "python":
		fmt.Println("python") // fallthrough가 있어서 출력
		fallthrough
	case "ruby":
		fmt.Println("ruby") // fallthrough가 있어서 출력
		fallthrough
	case "c":
		fmt.Println("c") // fallthrough가 있어서 출력
		// 마지막에는 fallthrough를 사용할 필요는 없음.
	}
}
