package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
	원자적 연산 사용하기
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			//data += 1 // 원자성을 보장할 수 없음.
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//data -= 1 // 원자성을 보장할 수 없음.
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
