package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"sync"
	"time"
	_ "time"
)

var (
	maxNumOfWorkers    = 1
	toBeHandleTaskChan = make(chan string, maxNumOfWorkers)
	runnableTask       = make(map[string]struct{})
	lock               = &sync.Mutex{}
)

func main() {
	forever := make(chan struct{})

	go func() {
		for {
			r := rand.Intn(2)

			task := "B"
			if r == 0 {
				task = "A"
			}

			func() {
				lock.Lock()
				defer lock.Unlock()

				_, exists := runnableTask[task]
				if exists {
					// 429 Too many requests
					fmt.Printf("[%s] 이미 수행 대기 중인 작업 요청\n", task)
					return
				}

				runnableTask[task] = struct{}{}
				fmt.Printf("[%s] 작업 할당\n", task)
				toBeHandleTaskChan <- task
			}()

			sleepDuration := time.Duration(rand.Intn(500)+1000) * time.Millisecond
			<-time.After(sleepDuration)
		}
	}()

	for i := 0; i < maxNumOfWorkers; i++ {
		id := i
		go func() {
			executedCount := 0
			for task := range toBeHandleTaskChan {
				func() {
					defer func() {
						lock.Lock()
						defer lock.Unlock()
						delete(runnableTask, task)
						fmt.Printf("[%d][%d-%s] 처리 완료\n", id, executedCount, task)
					}()

					if task == "A" {
						executeA()
					} else if task == "B" {
						executeB()
					}
					executedCount = executedCount + 1
				}()
			}
		}()
	}

	<-forever
}

func executeA() string {
	<-time.After(time.Duration(rand.Intn(3000)) * time.Millisecond)
	return "A 이벤트 처리"
}

func executeB() string {
	<-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond)
	return "B 이벤트 처리"
}
