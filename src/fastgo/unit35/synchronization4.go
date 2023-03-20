package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	읽기 쓰기 뮤텍스를 통해 정확한 값 읽기와 쓰기를 보장한다.
	???? 뭔가 코드가 많이 바뀜. 동작 안함.
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data = 0

	go func() {
		for i := 0; i < 3; i++ {
			mutex := sync.RWMutex{}
			mutex.Lock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex := sync.RWMutex{}
			rwMutex.RLock()
			fmt.Println("read1 : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex := sync.RWMutex{}
			rwMutex.RLock()
			fmt.Println("read2 : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
