package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	함수를 한 번만 실행하기 (Once)
*/

func helloWorld() {
	fmt.Println("Hello world!")
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("gorutine : ", n)

			once.Do(helloWorld) // 고루틴은 3개지만 hello 함수는 한 번만 실행한다.
			// 복잡한 반복문 초기화 등에 유용할 수 있다.
		}(i)
	}

	fmt.Scanln()
}
