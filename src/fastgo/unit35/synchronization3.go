package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	경합으로 인해 값이 원하는대로 읽기 쓰기 동작이 완벽하게 보장되지 않음.
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int = 0

	go func() {
		for i := 0; i < 3; i++ {
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
