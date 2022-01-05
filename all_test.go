package goworker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	Init("test", 10000, 100)

	waitGroup := new(sync.WaitGroup)

	fmt.Println("xxx")
	for xxx := 0; xxx < 10; xxx++ {
		go func() {
			for i := 0; i < 100000; i++ {

				waitGroup.Add(1)

				Push("test", func() {

					time.Sleep(time.Millisecond)

					waitGroup.Done()
				})

			}
		}()
	}

	fmt.Println("123ok")
	waitGroup.Wait()
	fmt.Println("ok")
	t.Log("ok")

}
